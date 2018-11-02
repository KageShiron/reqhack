package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestRestError(t *testing.T) {
	assert := assert.New(t)

	rec := httptest.NewRecorder()
	err := RestError(rec, 200, "foo")
	assert.Nil(err)
	assert.Equal(rec.Body.String(), `{"error":{"message":"foo","code":200}}`)
	assert.Equal(rec.Code, 200)

	rec = httptest.NewRecorder()
	err = RestError(rec, 404, "Not Found")
	assert.Nil(err)
	assert.Equal(rec.Body.String(), `{"error":{"message":"Not Found","code":404}}`)
	assert.Equal(rec.Code, 404)
}

func TestRestSucceed(t *testing.T) {
	assert := assert.New(t)

	rec := httptest.NewRecorder()
	err := RestSucceed(rec, 200, "foo")
	assert.Nil(err)
	assert.Equal(rec.Body.String(), `{"success":{"message":"foo","code":200}}`)
	assert.Equal(rec.Code, 200)

	rec = httptest.NewRecorder()
	err = RestSucceed(rec, 404, "Not Found")
	assert.Nil(err)
	assert.Equal(rec.Body.String(), `{"success":{"message":"Not Found","code":404}}`)
	assert.Equal(rec.Code, 404)
}
