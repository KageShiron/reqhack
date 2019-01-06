package infrastracture

import (
	"github.com/KageShiron/reqhack/server/domain"
)

// BinRepository is a bin manager
type BinRepository interface {
	Add(name string, secret string) (*domain.Bin, error)
	//Update(r *domain.Bin) error
	Get(name string, secret string) (*domain.Bin, error)
	GetWithoutSecret(name string) (*domain.Bin, error)
	GetByID(id int64, secret string) (bin *domain.Bin, err error)
}
