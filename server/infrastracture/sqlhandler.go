package infrastracture

import (
	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/jmoiron/sqlx"
)

// SQLHandler is a connection handler
type SQLHandler struct {
	Conn *sqlx.DB
}

// NewSQLHandler returns mysql connection
func NewSQLHandler() *SQLHandler {
	conn, err := sqlx.Open("mysql", "root:mysql@tcp(192.168.99.100:3306)/reqhack")
	if err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
