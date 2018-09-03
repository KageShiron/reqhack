package main

import (
	"fmt"
	"net/http"
	"time"
)

type Request struct {
	Time time.Time
	http.Request
}

/// Bin

type Bin struct {
	reqs *[]Request
}

func (bin Bin) WriteLog(request Request) (err error) {
	*bin.reqs = append(*bin.reqs, request);
	return nil
}

func (bin Bin) ReadLog(no int) (req *Request, err error) {
	if len(*bin.reqs) > no {
		return nil, fmt.Errorf("out of range")
	}
	return &(*bin.reqs)[no], nil
}

func (bin Bin) ReadLogs(index int, length int) (requests []Request, err error) {
	if len(*bin.reqs) > index+length {
		return nil, fmt.Errorf("out of range")
	}
	return (*bin.reqs)[index : index+length], nil
}
func (bin Bin) Length() int {
	return len(*(bin.reqs))
}

//////////

type BinManager struct {
	save map[string]*Bin
}

func NewBinManager() *BinManager {
	return &BinManager{save:make(map[string]*Bin)}
}

func (man BinManager) Create(binID string) *Bin {
	_, ok := man.save[binID]
	if ok {
		return nil
	}
	b := Bin{}
	man.save[binID] = new(Bin)
	return &b
}

func (man BinManager) Bin(binID string) *Bin {
	b, ok := man.save[binID]
	if !ok {
		return nil
	}
	return b
}
