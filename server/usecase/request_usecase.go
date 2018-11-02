package usecase

import (
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
)

// RequestUsecase is request manager
type RequestUsecase interface {
	Add(r *domain.Request) error
	Get(binID int64, id int64) (*domain.Request, error)
	GetRange(binID int64, start int64, length int64) ([]*domain.Request, error)
	Length(binID int64) (len int64, err error)
}

type requestUsecase struct {
	reqRepo infrastracture.RequestRepository
}

// NewRequestUsecase returns a new RequestUsecase
func NewRequestUsecase(reqRepo infrastracture.RequestRepository) RequestUsecase {
	return &requestUsecase{reqRepo: reqRepo}
}

// Add adds a HTTP Request log
func (b *requestUsecase) Add(r *domain.Request) error {
	return b.reqRepo.Add(r)
}

// Get returns a request
func (b *requestUsecase) Get(binID int64, id int64) (*domain.Request, error) {
	return b.reqRepo.Get(binID, id)
}

// GetRange returns Http Request Logs
func (b *requestUsecase) GetRange(binID int64, start int64, length int64) ([]*domain.Request, error) {
	return b.reqRepo.GetRange(binID, start, length)
}

// Length returns the length of the log records
func (b *requestUsecase) Length(binID int64) (len int64, err error) {
	return b.reqRepo.Length(binID)
}
