package infrastracture

import (
	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/jmoiron/sqlx"
	"os"
)

// SQLHandler is a connection handler
type SQLHandler struct {
	Conn *sqlx.DB
}

// NewSQLHandler returns mysql connection
func NewSQLHandler() *SQLHandler {
	src := os.Getenv("DATA_SOURCE_NAME")
	conn, err := sqlx.Open("mysql", src)
	if err != nil {
		panic(err.Error())
	}
	if err = conn.Ping(); err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
