package mysql

import (
	"fmt"
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
	"regexp"
)

var labelRegexp = regexp.MustCompile(`^([0-9a-zA-Z][0-9a-zA-Z\-]*[0-9a-zA-Z]|[0-9a-zA-Z])$`)
var binNameRegexp = regexp.MustCompile(`[0-9a-zA-Z][0-9a-zA-Z\-][0-9a-zA-Z]$`)

// mysqlBinRepository
type mysqlBinRepository struct {
	infrastracture.SQLHandler
}

// NewMysqlBinRepository returns new BinRepository
func NewMysqlBinRepository(handler infrastracture.SQLHandler) infrastracture.BinRepository {
	return &mysqlBinRepository{SQLHandler: handler}
}

func checkLabel( label string ) (error) {
	if labelRegexp.MatchString(label) {
		return nil
	}
	return fmt.Errorf("%s is invalid label" , label)
}

func getBinName( label string ) string {
	return binNameRegexp.FindString(label)
}
// Add adds new Bin
func (m *mysqlBinRepository) Add(label string, secret string) (*domain.Bin, error) {
	if err := checkLabel(label); err != nil {
		return nil, err
	}
	res, err := m.Conn.Exec("INSERT INTO `bin` (name,secret) VALUES (?,?)", label, secret)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &domain.Bin{ID: id, Name: label}, nil
}

// Get returns a bin
func (m *mysqlBinRepository) Get(label string, secret string) (bin *domain.Bin, err error) {
	name := getBinName(label)
	if name == "" {
		return nil, fmt.Errorf("invalid hostname %s",label)
	}
	bin = &domain.Bin{}
	print(label,name)
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE name=? AND secret=?", name, secret).StructScan(bin)
	return
}

// Get returns a bin
func (m *mysqlBinRepository) GetWithoutSecret(label string) (bin *domain.Bin, err error) {
	name := getBinName(label)
	if name == "" {
		return nil, fmt.Errorf("invalid hostname %s",label)
	}
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT name,name FROM bin WHERE name=?", name).StructScan(bin)
	return
}

// GetById returns a bin by id
func (m *mysqlBinRepository) GetByID(id int64, secret string) (bin *domain.Bin, err error) {
	bin = &domain.Bin{}
	err = m.Conn.QueryRowx("SELECT id,name FROM bin WHERE id=? AND secret=?", id, secret).StructScan(bin)
	return
}
