package usecase

import (
	"errors"
	"fmt"
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func createMockRequest(id int64, target string) *domain.Request {
	httpreq := httptest.NewRequest("GET", target, strings.NewReader("hogehoge"))
	req, _ := domain.NewRequest(time.Now(), httpreq)
	req.Bin = &domain.Bin{ID: id, Name: fmt.Sprintf("test%d", id)}
	return req
}

func TestReqAdd(t *testing.T) {
	mReqRepo := new(mocks.RequestRepository)
	mReq1 := createMockRequest(1, "http://example.com/foo")
	mReq2 := createMockRequest(2, "http://bar.example.com")

	mReqRepo.On("Add", mReq1).Return(nil)
	mReqRepo.On("Add", mReq2).Return(errors.New("error"))

	u := NewRequestUsecase(mReqRepo)
	err := u.Add(mReq1)
	assert.NoError(t, err)

	err = u.Add(mReq2)
	assert.Error(t, err)
}

func TestReqGet(t *testing.T) {
	mReqRepo := new(mocks.RequestRepository)
	mReq := createMockRequest(1, "http://example.com")

	mReqRepo.On("Get", int64(1), int64(1)).Return(mReq, nil)
	mReqRepo.On("Get", int64(1), int64(2)).Return(nil, errors.New("error"))

	u := NewRequestUsecase(mReqRepo)
	rr, err := u.Get(1, 1)
	assert.NoError(t, err)
	assert.EqualValues(t, mReq, rr)

	rr, err = u.Get(1, 2)
	assert.Error(t, err)
}

func TestReqGetRange(t *testing.T) {
	mReqRepo := new(mocks.RequestRepository)

	mReq1 := createMockRequest(1, "http://example.com/foo")
	mReq2 := createMockRequest(2, "http://bar.example.com")
	mReqs := []*domain.Request{mReq1, mReq2}

	mReqRepo.On("GetRange", int64(1), int64(1), int64(2)).Return(mReqs, nil)
	mReqRepo.On("GetRange", int64(1), int64(3), int64(2)).Return(nil, errors.New("error"))

	u := NewRequestUsecase(mReqRepo)
	rr, err := u.GetRange(1, 1, 2)
	assert.NoError(t, err)
	assert.EqualValues(t, mReqs, rr)

	rr, err = u.GetRange(1, 3, 2)
	assert.Error(t, err)
}
func TestReqLength(t *testing.T) {
	mReqRepo := new(mocks.RequestRepository)

	mReqRepo.On("Length", int64(1)).Return(int64(5), nil)
	mReqRepo.On("Length", int64(2)).Return(int64(0), errors.New("error"))

	u := NewRequestUsecase(mReqRepo)
	l, err := u.Length(1)
	assert.NoError(t, err)
	assert.EqualValues(t, l, 5)

	l, err = u.Length(2)
	assert.Error(t, err)
}
