package mysql

import (
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
)

// mysqlBinRepository
type mysqlBinRepository struct {
	infrastracture.SQLHandler
}

// NewMysqlBinRepository returns new BinRepository
func NewMysqlBinRepository(handler infrastracture.SQLHandler) infrastracture.BinRepository {
	return &mysqlBinRepository{SQLHandler: handler}
}

// Add adds new Bin
func (m *mysqlBinRepository) Add(name string) (*domain.Bin, error) {
	res, err := m.Conn.Exec("INSERT INTO `bin` (name) VALUES (?)", name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &domain.Bin{ID: id, Name: name}, nil
}

// Get returns a bin
func (m *mysqlBinRepository) Get(name string) (bin *domain.Bin, err error) {
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE name=?", name).StructScan(bin)
	return
}

// GetById returns a bin by id
func (m *mysqlBinRepository) GetByID(id int64) (bin *domain.Bin, err error) {
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE id=?", id).StructScan(bin)
	return
}
