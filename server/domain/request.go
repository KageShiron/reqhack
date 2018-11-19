package domain

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// A Request represents a log record of HTTP Request.
type Request struct {
	ID            int64           `json:"-"`
	Bin           *Bin            `json:"-"`
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

// A Bin represents a bin object
type Bin struct {
	ID   int64
	Name string
}

var realIPrand = "X-Reqhack-Real-IP-" + os.Getenv("REQHACK_RANDOM")
var baseHost = os.Getenv("REQHACK_BASEHOST")

// NewRequest return a Request object
func NewRequest(time time.Time, r *http.Request) (req *Request, err error) {
	body, err := ioutil.ReadAll(r.Body)
	// reassign body
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	r.ParseForm()
	//m, err := r.MultipartReader()
	println(realIPrand)
	println(r.Header.Get(realIPrand))
	ip := r.Header.Get(realIPrand)
	if ip != "" {
		r.Header.Del(realIPrand)
	} else {
		ip = r.RemoteAddr
	}

	pos := strings.LastIndex(r.Host,baseHost)
	prefix := "/v1/" + r.Host[:(pos-1)] + "/in"
	if !strings.HasPrefix(r.RequestURI,prefix){
		log.Fatal("Bad Prefix : " + r.RequestURI)
	}

	req = &Request{
		Time:       time,
		Method:     r.Method,
		Proto:      r.Proto,
		Header:     r.Header,
		Body:       body,
		Host:       r.Host,
		Form:       r.Form,
		PostForm:   r.PostForm,
		RemoteAddr: ip,
		RequestURI: strings.TrimPrefix(r.RequestURI,prefix),
	}
	return req, err
}

// A BinManager managements bins
type BinManager interface {
	Create(binID string) Bin
	Bin(binID string) Bin
}
