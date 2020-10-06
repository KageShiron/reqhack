package mysql

import (
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestReqGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	data := `{"body": "hoge", "form": {}, "host": "localhost:8081", "time": "2018-10-25T03:09:17.230635423+09:00", "header": {"Accept": ["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"], "Cookie": ["Phpstorm-299628a4=f7104557-4462-40ed-8761-8d438f268480"], "Connection": ["keep-alive"], "User-Agent": ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:62.0) Gecko/20100101 Firefox/62.0"], "Cache-Control": ["max-age=0"], "Accept-Encoding": ["gzip, deflate"], "Accept-Language": ["ja,en-US;q=0.7,en;q=0.3"], "Upgrade-Insecure-Requests": ["1"]}, "method": "GET", "postform": {}, "protocol": "HTTP/1.1", "remoteaddr": "127.0.0.1:54793", "requesturi": "/v1/nya/in/", "multipartform": null}`
	rows := sqlmock.NewRows([]string{"data"}).AddRow(data)
	query := "SELECT \\(data\\) FROM `request` WHERE bin=\\? AND id=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	sx := sqlx.NewDb(db, "sqlmock")
	sh := infrastracture.SQLHandler{Conn: sx}
	r := NewMysqlRequestRepository(sh)

	req, err := r.Get(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, req.ID, int64(0))
	assert.Nil(t, req.Bin)
	assert.Equal(t, req.RequestURI, "/v1/nya/in/")

	req, err = r.Get(1, 3)
	assert.Error(t, err)

	req, err = r.Get(2, 1)
	assert.Error(t, err)
}
func TestReqLength(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	query := "SELECT COUNT\\(\\*\\) FROM `request` WHERE bin=\\?"

	rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(5)
	mock.ExpectQuery(query).WillReturnRows(rows)

	sx := sqlx.NewDb(db, "sqlmock")
	sh := infrastracture.SQLHandler{Conn: sx}
	r := NewMysqlRequestRepository(sh)

	len, err := r.Length(1)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), len)
}

func TestReqAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()
	httpreq := httptest.NewRequest("GET", "http://example.com/hoge", strings.NewReader("hogehoge"))
	req, err := domain.NewRequest(time.Now(), httpreq)
	req.Bin = &domain.Bin{ID: 1, Name: "test"}

	query := "INSERT INTO `request` \\(bin,data\\) VALUES \\(\\?,\\?\\)"
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(3, 1))

	sx := sqlx.NewDb(db, "sqlmock")
	sh := infrastracture.SQLHandler{Conn: sx}
	r := NewMysqlRequestRepository(sh)

	err = r.Add(req)
	assert.NoError(t, err)
}
