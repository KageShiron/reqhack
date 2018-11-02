package mysql

import (
	"bytes"
	"encoding/json"
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
	"log"
)

// mysqlRequestRepository
type mysqlRequestRepository struct {
	infrastracture.SQLHandler
}

// NewMysqlRequestRepository returns new RequestRepository
func NewMysqlRequestRepository(handler infrastracture.SQLHandler) infrastracture.RequestRepository {
	return &mysqlRequestRepository{SQLHandler: handler}
}

// Add adds a HTTP Request log
func (m *mysqlRequestRepository) Add(r *domain.Request) (err error) {
	text, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := m.Conn.Exec("INSERT INTO `request` (bin,data) VALUES (?,?)", r.Bin.ID, text)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	r.ID = id
	return
}

// Get returns a request
func (m *mysqlRequestRepository) Get(binID int64, id int64) (*domain.Request, error) {
	row := m.Conn.QueryRow("SELECT (data) FROM `request` WHERE bin=? AND id=?", binID, id)
	var res string
	req := &domain.Request{}
	err := row.Scan(&res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes.NewBufferString(res).Bytes(), req)
	return req, err
}

// GetRange returns Http Request Logs
func (m *mysqlRequestRepository) GetRange(binID int64, start int64, length int64) ([]*domain.Request, error) {
	rows, err := m.Conn.Query("SELECT (data) FROM `request` WHERE bin=? AND id BETWEEN ? AND ?", binID, start, start+length-1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []byte

	reqs := []*domain.Request{}
	for rows.Next() {
		rows.Scan()
		req := &domain.Request{}
		reqs = append(reqs, req)
		err = rows.Scan(&res)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(res, req)
		if err != nil {
			return nil, err
		}
	}
	return reqs, nil
}

// Length returns the length of the log records
func (m *mysqlRequestRepository) Length(binID int64) (len int64, err error) {
	row := m.Conn.QueryRow("SELECT COUNT(*) FROM `request` WHERE bin=?", binID)
	err = row.Scan(&len)
	return
}
