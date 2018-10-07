package db

import (
	"database/sql"
	"github.com/lestrrat-go/test-mysqld"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	m.Run()
	//bin = MysqlBin{}
	//m.Run()
}

func TestNewRequest(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	hreq := httptest.NewRequest("GET", "http://example.com/hogehoge", strings.NewReader("foo"))
	req, err := NewRequest(now, hreq)
	assert.Nil(err)
	assert.Equal(req.Time, now)
	assert.Equal(req.Method, "GET")
	assert.Equal(req.Body, []byte("foo"))
	assert.Equal(req.RequestURI, "http://example.com/hogehoge")

	// TODO: formのテスト
}

func TestBin(t *testing.T) {
	assert := assert.New(t)

	r := []*Request{}
	bins := []Bin{MemoryBin{reqs: &r}}

	for _, bin := range bins {
		//初期化
		now := time.Now()
		httpReq := httptest.NewRequest("GET", "http://example.com/hogehoge", strings.NewReader("test"))
		httpReq2 := httptest.NewRequest("POST", "http://example.com/piyopiyo", strings.NewReader("x=0"))
		req1, err := NewRequest(now, httpReq)
		req2, err := NewRequest(now, httpReq2)

		///////////////////////////////////////
		// 0,1に対するReadLogは失敗する
		rReq1, err := bin.ReadLog(0)
		assert.Error(err)
		assert.Nil(rReq1)

		rReq2, err := bin.ReadLog(1)
		assert.Error(err)
		assert.Nil(rReq2)

		rReq3, err := bin.ReadLog(-1)
		assert.Error(err)
		assert.Nil(rReq3)

		length := bin.Length()
		assert.Equal(length, 0)

		logs1, err := bin.ReadLogs(0, 0)
		assert.Error(err)
		assert.Nil(logs1)

		logs2, err := bin.ReadLogs(0, 1)
		assert.Error(err)
		assert.Nil(logs2)

		logs3, err := bin.ReadLogs(-1, 1)
		assert.Error(err)
		assert.Nil(logs3)

		///////////////////////////////////////
		// 書き込みは成功する
		err = bin.WriteLog(req1)
		assert.Nil(err)

		// 0に対するReadLogは成功する
		// 1に対するReadLogは失敗する
		rReqA1, err := bin.ReadLog(0)
		assert.Nil(err)
		assert.Equal(rReqA1, req1)

		rReqA2, err := bin.ReadLog(1)
		assert.Error(err)
		assert.Nil(rReqA2)

		lengthA := bin.Length()
		assert.Equal(lengthA, 1)

		logsA1, err := bin.ReadLogs(0, 0)
		assert.Error(err)
		assert.Nil(logsA1)

		logsA2, err := bin.ReadLogs(-1, 1)
		assert.Error(err)
		assert.Nil(logsA2)

		logsA3, err := bin.ReadLogs(0, 1)
		assert.Nil(err)
		assert.Equal(len(logsA3), 1)
		assert.Equal(logsA3[0], req1)

		logsA4, err := bin.ReadLogs(0, 2)
		assert.Nil(err)
		assert.Equal(len(logsA4), 1)
		assert.Equal(logsA4[0], req1)

		logsA5, err := bin.ReadLogs(1, 2)
		assert.Nil(logsA5)
		assert.Error(err)

		///////////////////////////////////////
		// 書き込みは成功する
		err = bin.WriteLog(req2)
		assert.Nil(err)

		// 0に対するReadLogは成功する
		// 1に対するReadLogは失敗する
		rReqB1, err := bin.ReadLog(0)
		assert.Nil(err)
		assert.Equal(rReqB1, req1)

		rReqB2, err := bin.ReadLog(1)
		assert.Nil(err)
		assert.Equal(rReqB2, req2)

		lengthB := bin.Length()
		assert.Equal(lengthB, 2)

		logsB1, err := bin.ReadLogs(0, 0)
		assert.Error(err)
		assert.Nil(logsB1)

		logsB2, err := bin.ReadLogs(-1, 1)
		assert.Error(err)
		assert.Nil(logsB2)

		logsB3, err := bin.ReadLogs(0, 1)
		assert.Nil(err)
		assert.Equal(len(logsB3), 1)
		assert.Equal(logsB3[0], req1)

		logsB4, err := bin.ReadLogs(0, 2)
		assert.Nil(err)
		assert.Equal(len(logsB4), 2)
		assert.Equal(logsB4[0], req1)
		assert.Equal(logsB4[1], req2)

		logsB5, err := bin.ReadLogs(1, 2)
		assert.Nil(err)
		assert.Equal(len(logsB5), 1)
		assert.Equal(logsB5[0], req2)
	}
}

func TestBinManager(t *testing.T) {
	assert := assert.New(t)

	mysqld, err := mysqltest.NewMysqld(nil)
	if err != nil {
		log.Fatalf("Failed to start mysqld: %s", err)
	}
	defer mysqld.Stop()
	ds := mysqld.Datasource("test", "", "", 0)
	db, err := sql.Open("mysql", ds)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("CREATE TABLE `teset`"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("USE test;"); err != nil {
		t.Fatal(err)
	}
	mans := []BinManager{NewMemoryBinManager(), NewMysqlBinManager(ds)}

	now := time.Now()
	httpReq := httptest.NewRequest("GET", "http://example.com/hogehoge", strings.NewReader("test"))
	httpReq2 := httptest.NewRequest("POST", "http://example.com/piyopiyo", strings.NewReader("x=0"))
	req1, _ := NewRequest(now, httpReq)
	req2, _ := NewRequest(now, httpReq2)

	for _, man := range mans {
		b1 := man.Create("hoge")
		assert.NotNil(b1)
		b1.WriteLog(req1)
		b := man.Create("hoge")
		assert.Nil(b)

		b2 := man.Create("Hoge")
		assert.NotNil(b2)
		_, err := b2.ReadLog(0)
		assert.Error(err)
		b2.WriteLog(req2)
		b = man.Create("Hoge")
		assert.Nil(b)

		c1 := man.Bin("hoge")
		assert.Equal(c1, b1)
		r1, err := c1.ReadLog(0)
		assert.Nil(err)
		assert.Equal(r1, req1)

		c2 := man.Bin("Hoge")
		r2, err := c2.ReadLog(0)
		assert.Nil(err)
		assert.Equal(r2, req2)
		assert.Equal(c2, b2)

		b = man.Bin("HOGE")
		assert.Nil(b)
	}
}
