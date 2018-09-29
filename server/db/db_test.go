package db

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var bin Bin

func TestMain(m *testing.M) {
	r := []*Request{}
	bin = MemoryBin{reqs: &r}
	m.Run()
	//bin = MysqlBin{}
	//m.Run()
}

func TestNewRequest(t *testing.T) {
	assert := assert.New(t)

	now := time.Now()
	hreq := httptest.NewRequest("GET", "http://example.com/hogehoge", strings.NewReader("foo"))
	req, err := NewRequest(now, hreq)
	assert.Nil(err)
	assert.Equal(req.Time, now)
	assert.Equal(req.Method, "GET")
	assert.Equal(req.Body, []byte("foo"))
	assert.Equal(req.RequestURI, "http://example.com/hogehoge")

	// TODO: formのテスト
}

func TestBin(t *testing.T) {
	assert := assert.New(t)

	//初期化
	now := time.Now()
	httpReq := httptest.NewRequest("GET", "http://example.com/hogehoge", strings.NewReader("test"))
	httpReq2 := httptest.NewRequest("POST", "http://example.com/piyopiyo", strings.NewReader("x=0"))
	req1, err := NewRequest(now, httpReq)
	req2, err := NewRequest(now, httpReq2)

	///////////////////////////////////////
	// 0,1に対するReadLogは失敗する
	rReq1, err := bin.ReadLog(0)
	assert.Error(err)
	assert.Nil(rReq1)

	rReq2, err := bin.ReadLog(1)
	assert.Error(err)
	assert.Nil(rReq2)

	rReq3, err := bin.ReadLog(-1)
	assert.Error(err)
	assert.Nil(rReq3)

	length := bin.Length()
	assert.Equal(length, 0)

	logs1, err := bin.ReadLogs(0, 0)
	assert.Error(err)
	assert.Nil(logs1)

	logs2, err := bin.ReadLogs(0, 1)
	assert.Error(err)
	assert.Nil(logs2)

	logs3, err := bin.ReadLogs(-1, 1)
	assert.Error(err)
	assert.Nil(logs3)

	///////////////////////////////////////
	// 書き込みは成功する
	err = bin.WriteLog(req1)
	assert.Nil(err)

	// 0に対するReadLogは成功する
	// 1に対するReadLogは失敗する
	rReqA1, err := bin.ReadLog(0)
	assert.Nil(err)
	assert.Equal(rReqA1, req1)

	rReqA2, err := bin.ReadLog(1)
	assert.Error(err)
	assert.Nil(rReqA2)

	lengthA := bin.Length()
	assert.Equal(lengthA, 1)

	logsA1, err := bin.ReadLogs(0, 0)
	assert.Error(err)
	assert.Nil(logsA1)

	logsA2, err := bin.ReadLogs(-1, 1)
	assert.Error(err)
	assert.Nil(logsA2)

	logsA3, err := bin.ReadLogs(0, 1)
	assert.Nil(err)
	assert.Equal(len(logsA3), 1)
	assert.Equal(logsA3[0], req1)

	logsA4, err := bin.ReadLogs(0, 2)
	assert.Nil(err)
	assert.Equal(len(logsA4), 1)
	assert.Equal(logsA4[0], req1)

	logsA5, err := bin.ReadLogs(1, 2)
	assert.Nil(logsA5)
	assert.Error(err)

	///////////////////////////////////////
	// 書き込みは成功する
	err = bin.WriteLog(req2)
	assert.Nil(err)

	// 0に対するReadLogは成功する
	// 1に対するReadLogは失敗する
	rReqB1, err := bin.ReadLog(0)
	assert.Nil(err)
	assert.Equal(rReqB1, req1)

	rReqB2, err := bin.ReadLog(1)
	assert.Nil(err)
	assert.Equal(rReqB2, req2)

	lengthB := bin.Length()
	assert.Equal(lengthB, 2)

	logsB1, err := bin.ReadLogs(0, 0)
	assert.Error(err)
	assert.Nil(logsB1)

	logsB2, err := bin.ReadLogs(-1, 1)
	assert.Error(err)
	assert.Nil(logsB2)

	logsB3, err := bin.ReadLogs(0, 1)
	assert.Nil(err)
	assert.Equal(len(logsB3), 1)
	assert.Equal(logsB3[0], req1)

	logsB4, err := bin.ReadLogs(0, 2)
	assert.Nil(err)
	assert.Equal(len(logsB4), 2)
	assert.Equal(logsB4[0], req1)
	assert.Equal(logsB4[1], req2)

	logsB5, err := bin.ReadLogs(1, 2)
	assert.Nil(err)
	assert.Equal(len(logsB5), 1)
	assert.Equal(logsB5[0], req2)
}
