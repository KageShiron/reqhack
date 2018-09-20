package db

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

// A Request represents a log record of HTTP Request.
type Request struct {
	Time          time.Time       `json:"time"`
	Method        string          `json:"method"`
	Proto         string          `json:"protocol"`
	Header        http.Header     `json:"header"`
	Body          []byte          `json:"body"`
	Host          string          `json:"host"`
	Form          url.Values      `json:"form"`
	PostForm      url.Values      `json:"postform"`
	MultipartForm *multipart.Form `json:"multipartform"`
	RemoteAddr    string          `json:"remoteaddr"`
	RequestURI    string          `json:"requesturi"`
}

// NewRequest return a Request object
func NewRequest(time time.Time, r *http.Request) (req *Request, err error) {
	body, err := ioutil.ReadAll(r.Body)
	// reassign body
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	r.ParseForm()
	//m, err := r.MultipartReader()

	req = &Request{
		Time:       time,
		Method:     r.Method,
		Proto:      r.Proto,
		Header:     r.Header,
		Body:       body,
		Host:       r.Host,
		Form:       r.Form,
		PostForm:   r.PostForm,
		RemoteAddr: r.RemoteAddr,
		RequestURI: r.RequestURI,
	}
	return req, err
}

///////////////////////////////

// A Bin represents a requests bin box.
type Bin struct {
	reqs *[]*Request
}

// WriteLog writes a HTTP Request log
func (bin Bin) WriteLog(request *Request) (err error) {
	*bin.reqs = append(*bin.reqs, request)
	return nil
}

// ReadLog returns a log
func (bin Bin) ReadLog(no int) (req *Request, err error) {
	if len(*bin.reqs) <= no {
		return nil, fmt.Errorf("out of range")
	}
	return (*bin.reqs)[no], nil
}

// ReadLogs returns Http Request Logs
func (bin Bin) ReadLogs(index int, length int) (requests []*Request, err error) {
	if len(*bin.reqs) > index+length {
		return nil, fmt.Errorf("out of range")
	}
	return (*bin.reqs)[index : index+length], nil
}

// Length returns the length of the log records
func (bin Bin) Length() int {
	return len(*(bin.reqs))
}

//////////////////

// A BinManager managements bins
type BinManager struct {
	save map[string]*Bin
}

// NewBinManager returns a pointer of a new BinManager
func NewBinManager() *BinManager {
	return &BinManager{save: make(map[string]*Bin)}
}

// Create return new Bin object.
func (man BinManager) Create(binID string) *Bin {
	_, ok := man.save[binID]
	if ok {
		return nil
	}
	r := []*Request{}
	b := Bin{reqs: &r}
	man.save[binID] = &b
	return &b
}

// Bin returns bin which binID pointing
func (man BinManager) Bin(binID string) *Bin {
	b, ok := man.save[binID]
	if !ok {
		return nil
	}
	return b
}
