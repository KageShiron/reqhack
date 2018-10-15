package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	// mySql driver
	_ "github.com/go-sql-driver/mysql"
)

// A MysqlBin represents a requests bin box on mysql.
type MysqlBin struct {
	manager *MysqlBinManager
	id      int64
}

// MysqlBinManager managements bins on memory
type MysqlBinManager struct {
	db *sql.DB
}

//////////

// WriteLog writes a HTTP Request log
func (bin *MysqlBin) WriteLog(request *Request) (err error) {
	mar, err := json.Marshal(request)
	//mar, err := json.Unmarshal(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = bin.manager.db.Exec("INSERT INTO `request` (bin,data) VALUES (?,?)", bin.id, mar)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

// ReadLog returns a log
func (bin *MysqlBin) ReadLog(no int) (req *Request, err error) {
	row := bin.manager.db.QueryRow("SELECT (data) FROM `request` WHERE bin=? AND id=?", bin.id, no)
	var res []byte
	req = &Request{}
	err = row.Scan(&res)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, req)
	return
}

// SetResponser sets a responser
func (bin *MysqlBin) SetResponser(data Responser) {
	mar, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = bin.manager.db.Exec("INSERT INTO `response` (bin,data) VALUES (?,?)", bin.id, mar)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

// GetResponser gets a responser
func (bin *MysqlBin) GetResponser(reqURL string) (responser *Responser, err error) {
	u, err := url.Parse(reqURL)
	if err != nil {
		return nil, fmt.Errorf("bad url")
	}
	row := bin.manager.db.QueryRow("SELECT (data) FROM `responser` WHERE bin=? AND path = BINARY ?", bin.id, u.Path)
	var res []byte
	responser = &Responser{}
	err = row.Scan(&res)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, responser)
	return
}

// ReadLogs returns Http Request Logs
func (bin *MysqlBin) ReadLogs(index int, length int) (requests []*Request, err error) {
	rows, err := bin.manager.db.Query("SELECT (data) FROM `request` WHERE bin=? AND id BETWEEN ? AND ?", bin.id, index, index+length-1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []byte

	requests = []*Request{}
	for rows.Next() {
		rows.Scan()
		req := &Request{}
		requests = append(requests, req)
		err = rows.Scan(&res)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(res, req)
		if err != nil {
			return nil, err
		}
	}
	return requests, nil
}

// Length returns the length of the log records
func (bin *MysqlBin) Length() (len int) {
	row := bin.manager.db.QueryRow("COUNT( SELECT (data) FROM `request` WHERE bin=? )", bin.id)
	row.Scan(&len)
	return
}

//////////

// NewMysqlBinManager returns a pointer of a new BinManager for mysql
func NewMysqlBinManager(dataSourceName string) *MysqlBinManager {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &MysqlBinManager{db: db}
}

// Create return new Bin object.
func (man *MysqlBinManager) Create(binID string) Bin {
	res, err := man.db.Exec("INSERT INTO `bin` (name) VALUES (?)", binID)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &MysqlBin{manager: man, id: id}
}

// Bin returns bin which binID pointing
func (man *MysqlBinManager) Bin(binID string) Bin {
	row := man.db.QueryRow("SELECT (id) FROM bin WHERE name=?", binID)
	if row == nil {
		return nil
	}
	var id int64
	row.Scan(&id)
	return &MysqlBin{manager: man, id: id}
}
