package infrastracture

import (
	"github.com/KageShiron/reqhack/server/domain"
)

// BinRepository is a bin manager
type BinRepository interface {
	Add(label string, secret string) (*domain.Bin, error)
	//Update(r *domain.Bin) error
	Get(label string, secret string) (*domain.Bin, error)
	GetWithoutSecret(label string) (*domain.Bin, error)
	GetByID(id int64, secret string) (bin *domain.Bin, err error)
}
