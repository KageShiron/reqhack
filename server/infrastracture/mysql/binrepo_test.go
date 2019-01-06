package mysql

import (
	"github.com/KageShiron/reqhack/server/infrastracture"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestBinGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").AddRow(2, "bar")
	query := "SELECT \\* FROM bin WHERE name=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	sx := sqlx.NewDb(db, "sqlmock")
	sh := infrastracture.SQLHandler{Conn: sx}
	b := NewMysqlBinRepository(sh)

	req, err := b.Get("foo", "")
	assert.NoError(t, err)
	assert.Equal(t, req.Name, "foo")
	assert.Equal(t, req.ID, int64(1))

	req, err = b.Get("nothing", "")
	assert.Error(t, err)
}

func TestBinGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "foo").AddRow(2, "bar")
	query := "SELECT \\* FROM bin WHERE id=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	sx := sqlx.NewDb(db, "sqlmock")
	sh := infrastracture.SQLHandler{Conn: sx}
	b := NewMysqlBinRepository(sh)

	req, err := b.GetByID(1, "")
	assert.NoError(t, err)
	assert.Equal(t, req.Name, "foo")
	assert.Equal(t, req.ID, int64(1))

	req, err = b.GetByID(9, "")
	assert.Error(t, err)
}

func TestBinAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	defer db.Close()

	query := "INSERT INTO `bin` \\(name\\) VALUES \\(\\?\\)"
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(3, 1))

	sx := sqlx.NewDb(db, "sqlmock")
	sh := infrastracture.SQLHandler{Conn: sx}
	b := NewMysqlBinRepository(sh)

	req, err := b.Add("hoge", "")
	assert.NoError(t, err)
	assert.Equal(t, req.Name, "hoge")
	assert.Equal(t, req.ID, int64(3))

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewErrorResult(errors.New("test")))
	req, err = b.Add("hoge", "")
	assert.Error(t, err)
}
