package infrastracture

import (
	"github.com/KageShiron/reqhack/server/domain"
)

// RequestRepository is request manager
type RequestRepository interface {
	Add(r *domain.Request) (err error)
	Get(binID int64, id int64) (*domain.Request, error)
	GetRange(binID int64, start int64, length int64) ([]*domain.Request, error)
	Length(binID int64) (len int64, err error)
}
