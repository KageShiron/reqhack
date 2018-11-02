package usecase

import (
	"errors"
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinGet(t *testing.T) {
	mBinRepo := new(mocks.BinRepository)
	mBin := &domain.Bin{Name: "foo", ID: 1}

	mBinRepo.On("Get", "foo").Return(mBin, nil)
	mBinRepo.On("Get", "bar").Return(nil, errors.New("error"))

	u := NewBinUsecase(mBinRepo)
	rb, err := u.Get("foo")
	assert.NoError(t, err)
	assert.EqualValues(t, mBin, rb)

	rb, err = u.Get("bar")
	assert.Error(t, err)
}

func TestBinGetByID(t *testing.T) {
	mBinRepo := new(mocks.BinRepository)
	mBin := &domain.Bin{Name: "foo", ID: 1}

	mBinRepo.On("GetByID", int64(1)).Return(mBin, nil)
	mBinRepo.On("GetByID", int64(2)).Return(nil, errors.New("error"))

	u := NewBinUsecase(mBinRepo)
	rb, err := u.GetByID(1)
	assert.NoError(t, err)
	assert.EqualValues(t, mBin, rb)

	rb, err = u.GetByID(2)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	mBinRepo := new(mocks.BinRepository)
	mBin := &domain.Bin{Name: "foo", ID: 1}

	mBinRepo.On("Add", "foo").Return(mBin, nil)
	mBinRepo.On("Add", "bar").Return(nil, errors.New("error"))

	u := NewBinUsecase(mBinRepo)
	rb, err := u.Add("foo")
	assert.NoError(t, err)
	assert.EqualValues(t, mBin, rb)

	rb, err = u.Add("bar")
	assert.Error(t, err)
}
