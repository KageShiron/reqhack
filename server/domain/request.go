package domain

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// A Request represents a log record of HTTP Request.
type Request struct {
	ID            int64           `json:"-"`
	Bin           *Bin            `json:"-"`
	Time          time.Time       `json:"time"`
	ServerPort    int             `json:"server_port"`
	UserPort      int             `json:"user_port"`
	Scheme        string          `json:"scheme"`
	RawRequest    string          `json:"rawrequest"`
	Method        string          `json:"method"`
	Proto         string          `json:"protocol"`
	Header        http.Header     `json:"header"`
	Body          []byte          `json:"body"`
	BodyLength    int             `json:"body_length"`
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

var reqhackRandom = os.Getenv("REQHACK_RANDOM")
var realIP = "X-Reqhack-Real-IP-" + reqhackRandom
var realUserPort = "X-Reqhack-Real-UserPort-" + reqhackRandom
var realServerPort = "X-Reqhack-Real-ServerPort-" + reqhackRandom
var realScheme = "X-Reqhack-Real-Scheme-" + reqhackRandom
var realRequest = "X-Reqhack-Real-Request-" + reqhackRandom
var baseHost = os.Getenv("REQHACK_BASEHOST")

// NewRequest return a Request object
func NewRequest(time time.Time, r *http.Request) (req *Request, err error) {
	body, err := ioutil.ReadAll(r.Body)
	// reassign body
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	r.ParseForm()
	//m, err := r.MultipartReader()
	ip := r.Header.Get(realIP)
	userPort, _ := strconv.Atoi(r.Header.Get(realUserPort))
	serverPort, _ := strconv.Atoi(r.Header.Get(realServerPort))
	scheme := r.Header.Get(realScheme)
	request, _ := base64.StdEncoding.DecodeString(r.Header.Get(realRequest))
	prefix := ""
	r.Header.Del(realIP)
	r.Header.Del(realServerPort)
	r.Header.Del(realUserPort)
	r.Header.Del(realScheme)
	r.Header.Del(realRequest)
	if ip != "" {
		pos := strings.LastIndex(r.Host, baseHost)
		prefix = "/v1/" + r.Host[:(pos-1)] + "/in"
		if !strings.HasPrefix(r.RequestURI, prefix) {
			log.Fatal("Bad Prefix : " + r.RequestURI)

		}
	} else {
		ip = r.RemoteAddr
	}

	req = &Request{
		Time:       time,
		ServerPort: serverPort,
		UserPort:   userPort,
		Scheme:     scheme,
		RawRequest: string(request),
		Method:     r.Method,
		Proto:      r.Proto,
		Header:     r.Header,
		Body:       body,
		BodyLength: len(body),
		Host:       r.Host,
		Form:       r.Form,
		PostForm:   r.PostForm,
		RemoteAddr: ip,
		RequestURI: strings.TrimPrefix(r.RequestURI, prefix),
	}
	return req, err
}

// A BinManager managements bins
type BinManager interface {
	Create(binID string) Bin
	Bin(binID string) Bin
}
