package usecase

import (
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
)

// BinUsecase represents Bin's usecases
type BinUsecase interface {
	Add(name string) (*domain.Bin, error)
	Get(name string) (*domain.Bin, error)
	GetByID(id int64) (*domain.Bin, error)
}

type binUsecase struct {
	binRepo infrastracture.BinRepository
}

// NewBinUsecase returns new BinUsecase
func NewBinUsecase(binRepo infrastracture.BinRepository) BinUsecase {
	return &binUsecase{binRepo: binRepo}
}

// Add a bin
func (b *binUsecase) Add(name string) (*domain.Bin, error) {
	return b.binRepo.Add(name)
}

// Get returns a bin
func (b *binUsecase) Get(name string) (*domain.Bin, error) {
	return b.binRepo.Get(name)
}

// GetById return a bin by id
func (b *binUsecase) GetByID(id int64) (*domain.Bin, error) {
	return b.binRepo.GetByID(id)
}
