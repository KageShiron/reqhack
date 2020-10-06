package controller

import (
	"errors"
	"fmt"
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/usecase/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRequestIn(t *testing.T) {
	now := time.Now()
	NowFunc = func() time.Time {
		return now
	}

	bin := &mocks.BinUsecase{}
	bin.On("Get", "hoge").Return(nil, nil)
	bin.On("Get", "hoge2").Return(nil, errors.New("error"))
	req := &mocks.RequestUsecase{}
	req.On("Add", mock.AnythingOfType("*domain.Request")).Return(nil, nil)

	con := NewRequestController(req, bin)
	rec := httptest.NewRecorder()
	rt := mux.NewRouter()
	rt.Path("/v1/{name}/{_:in(?:/.*|$)}").HandlerFunc(con.In)

	r := httptest.NewRequest("POST", "http://example.com/v1/hoge/in", strings.NewReader("hoge"))
	rw, _ := domain.NewRequest(NowFunc(), r)
	rt.ServeHTTP(rec, r)
	assert.Equal(t, rec.Code, 200)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertCalled(t, "Add", rw)

	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	r2 := httptest.NewRequest("POST", "http://example.com/v1/hoge2/in", strings.NewReader("hoge"))
	rt.ServeHTTP(rec, r2)
	assert.Equal(t, rec.Code, 404)
	bin.AssertCalled(t, "Get", "hoge2")
	req.AssertNotCalled(t, "Add", mock.Anything)

	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	r3 := httptest.NewRequest("POST", "http://example.com/v1/hoge/in/test?hoge", strings.NewReader("hoge"))
	rw3, _ := domain.NewRequest(NowFunc(), r3)
	rt.ServeHTTP(rec, r3)
	assert.Equal(t, rec.Code, 200)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertCalled(t, "Add", rw3)
}

func TestRequestItems(t *testing.T) {
	now := time.Now()
	NowFunc = func() time.Time {
		return now
	}
	r1 := httptest.NewRequest("GET", "http://example.com/v1/hoge/items", strings.NewReader("hoge"))
	r2 := httptest.NewRequest("GET", "http://example.com/v1/hoge/items/", strings.NewReader("hoge"))
	r3 := httptest.NewRequest("POST", "http://example.com/v1/hoge/items", strings.NewReader("hoge"))
	r4 := httptest.NewRequest("GET", "http://example.com/v1/hoge/items/0", strings.NewReader("hoge"))
	rw1, _ := domain.NewRequest(NowFunc(), r1)
	rw2, _ := domain.NewRequest(NowFunc(), r2)
	rw3, _ := domain.NewRequest(NowFunc(), r3)
	rw4, _ := domain.NewRequest(NowFunc(), r4)

	bin := &mocks.BinUsecase{}
	b := &domain.Bin{ID: 1, Name: "hoge"}
	bin.On("Get", "hoge").Return(b, nil)
	bin.On("Get", "hoge2").Return(nil, errors.New("error"))
	req := &mocks.RequestUsecase{}
	req.On("GetRange", b.ID, int64(0), int64(100)).Return([]*domain.Request{rw1, rw2, rw3, rw4}, nil)
	req.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

	con := NewRequestController(req, bin)
	rt := mux.NewRouter()
	rt.HandleFunc("/v1/{name}/items/{num:[0-9]*(?:\\/)?}", con.Items).Methods("GET")
	rt.HandleFunc("/v1/{name}/items", con.Items).Methods("GET")

	// GET /v1/hoge/items
	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, r1)
	assert.Equal(t, rec.Code, 200)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertCalled(t, "GetRange", b.ID, int64(0), int64(100))

	// GET /v1/hoge/items/
	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	rt.ServeHTTP(rec, r2)
	assert.Equal(t, rec.Code, 200)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertCalled(t, "GetRange", b.ID, int64(0), int64(100))

	// POST /v1/hoge/items/ not accepted
	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	rt.ServeHTTP(rec, r3)
	assert.Equal(t, rec.Code, 405)
	bin.AssertNotCalled(t, "Get", "hoge")
	req.AssertNotCalled(t, "GetRange", b.ID, int64(0), int64(100))

	// GET /v1/hoge/items/0 don't call GetRange
	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	rt.ServeHTTP(rec, r4)
	assert.Equal(t, rec.Code, 200)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertNotCalled(t, "GetRange", b.ID, int64(0), int64(100))
}

func TestRequestItem(t *testing.T) {
	now := time.Now()
	NowFunc = func() time.Time {
		return now
	}
	r1 := httptest.NewRequest("GET", "http://example.com/v1/hoge/items/1", strings.NewReader("hoge"))
	r2 := httptest.NewRequest("GET", "http://example.com/v1/hoge/items/invalid", strings.NewReader("hoge"))
	r3 := httptest.NewRequest("POST", "http://example.com/v1/hoge/items/2", strings.NewReader("hoge"))
	r4 := httptest.NewRequest("GET", "http://example.com/v1/hoge/items/2", strings.NewReader("hoge"))
	rw1, _ := domain.NewRequest(NowFunc(), r1)

	bin := &mocks.BinUsecase{}
	b := &domain.Bin{ID: 1, Name: "hoge"}
	bin.On("Get", "hoge").Return(b, nil)
	bin.On("Get", "hoge2").Return(nil, errors.New("error"))
	req := &mocks.RequestUsecase{}
	req.On("Get", int64(1), int64(1)).Return(rw1, nil)
	req.On("Get", int64(1), int64(2)).Return(nil, errors.New("error"))

	con := NewRequestController(req, bin)
	rt := mux.NewRouter()
	rt.HandleFunc("/v1/{name}/items/{num:[0-9]*(?:\\/)?}", con.Items).Methods("GET")
	rt.HandleFunc("/v1/{name}/items", con.Items).Methods("GET")

	// GET /v1/hoge/items/1
	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, r1)
	assert.Equal(t, rec.Code, 200)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertCalled(t, "Get", int64(1), int64(1))
	timeStr, _ := NowFunc().MarshalJSON()
	assert.Equal(t, rec.Body.String(), fmt.Sprintf(`{"time":%s,"method":"GET","protocol":"HTTP/1.1","header":{},"body":"aG9nZQ==","host":"example.com","form":{},"postform":{},"multipartform":null,"remoteaddr":"192.0.2.1:1234","requesturi":"http://example.com/v1/hoge/items/1"}
`, timeStr))

	// GET /v1/hoge/items/invalid
	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	rt.ServeHTTP(rec, r2)
	assert.Equal(t, rec.Code, 404)
	bin.AssertNotCalled(t, "Get", "hoge")
	req.AssertNotCalled(t, "Get")

	// POST /v1/hoge/items/000 not accepted
	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	rt.ServeHTTP(rec, r3)
	assert.Equal(t, rec.Code, 405)

	// GET /v1/hoge/items/2 don't call GetRange
	rec = httptest.NewRecorder()
	bin.Calls = nil
	req.Calls = nil
	rt.ServeHTTP(rec, r4)
	assert.Equal(t, rec.Code, 404)
	bin.AssertCalled(t, "Get", "hoge")
	req.AssertCalled(t, "Get", int64(1), int64(2))
	fail := `{"error":{"message":"Log \"#2\" not found","code":404}}`
	assert.Equal(t, rec.Body.String(), fail)
}
