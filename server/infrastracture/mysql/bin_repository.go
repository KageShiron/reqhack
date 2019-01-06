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
func (m *mysqlBinRepository) Add(name string, secret string) (*domain.Bin, error) {
	res, err := m.Conn.Exec("INSERT INTO `bin` (name,secret) VALUES (?,?)", name, secret)
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
func (m *mysqlBinRepository) Get(name string, secret string) (bin *domain.Bin, err error) {
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE name=? AND secret=?", name, secret).StructScan(bin)
	return
}

// Get returns a bin
func (m *mysqlBinRepository) GetWithoutSecret(name string) (bin *domain.Bin, err error) {
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE name=?", name).StructScan(bin)
	return
}

// GetById returns a bin by id
func (m *mysqlBinRepository) GetByID(id int64, secret string) (bin *domain.Bin, err error) {
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE id=? AND secret=?", id, secret).StructScan(bin)
	return
}
