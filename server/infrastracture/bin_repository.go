package infrastracture

import (
	"github.com/KageShiron/reqhack/server/domain"
)

// BinRepository is a bin manager
type BinRepository interface {
	Add(name string) (*domain.Bin, error)
	//Update(r *domain.Bin) error
	Get(name string) (*domain.Bin, error)
	GetByID(id int64) (bin *domain.Bin, err error)
}
