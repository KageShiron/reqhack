package controller

import (
	mocks2 "github.com/KageShiron/reqhack/server/usecase/mocks"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBinCreate(t *testing.T) {
	bin := &mocks2.BinUsecase{}
	bin.On("Add", "hoge").Return(nil, nil).Once()
	bin.On("Add", "hoge2").Return(nil, errors.New("error")).Once()
	con := NewBinController(bin)

	rec := httptest.NewRecorder()
	rt := mux.NewRouter()
	rt.Methods("POST").Path("/v1/{name}/{_:create/?}").HandlerFunc(con.Create)
	r := httptest.NewRequest("POST", "http://example.com/v1/hoge/create", strings.NewReader("hoge"))
	rt.ServeHTTP(rec, r)
	suc := `{"success":{"message":"Created hoge bin","code":200}}`
	assert.Equal(t, suc, rec.Body.String())

	rec = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "http://example.com/v1/hoge2/create", strings.NewReader("hoge"))
	rt.ServeHTTP(rec, r)
	fail := `{"error":{"message":"Bin hoge2 already exists.","code":500}}`
	assert.Equal(t, fail, rec.Body.String())

	//rec.Body.Reset()
	//r = httptest.NewRequest("POST", "http://example.com/v1/ /create", strings.NewReader("hoge"))
	//rt.ServeHTTP(rec, r)
	//fail = `{"error":{"message":"Invalid name","code":400}}`
	//assert.Equal(t, fail, rec.Body.String())
}
