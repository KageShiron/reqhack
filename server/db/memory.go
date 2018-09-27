package db

import "fmt"

// MemoryBin represents a requests bin box on memory.
type MemoryBin struct {
	reqs *[]*Request
}

// WriteLog writes a HTTP Request log
func (bin MemoryBin) WriteLog(request *Request) (err error) {
	*bin.reqs = append(*bin.reqs, request)
	return nil
}

// ReadLog returns a log
func (bin MemoryBin) ReadLog(no int) (req *Request, err error) {
	if len(*bin.reqs) <= no {
		return nil, fmt.Errorf("out of range")
	}
	return (*bin.reqs)[no], nil
}

// ReadLogs returns Http Request Logs
func (bin MemoryBin) ReadLogs(index int, length int) (requests []*Request, err error) {
	loglength := len(*bin.reqs)
	if loglength < index {
		return nil, fmt.Errorf("out of range")
	}
	if index+length > loglength {
		return (*bin.reqs)[index:loglength], nil
	}

	return (*bin.reqs)[index:(index + length)], nil
}

// Length returns the length of the log records
func (bin MemoryBin) Length() int {
	return len(*(bin.reqs))
}

/////////////////////////

// MemoryBinManager managements bins on memory
type MemoryBinManager struct {
	save map[string]*MemoryBin
}

// NewMemoryBinManager returns a pointer of a new BinManager on memory
func NewMemoryBinManager() *MemoryBinManager {
	return &MemoryBinManager{save: make(map[string]*MemoryBin)}
}

// Create return new Bin object.
func (man *MemoryBinManager) Create(binID string) *MemoryBin {
	_, ok := man.save[binID]
	if ok {
		return nil
	}
	r := []*Request{}
	b := MemoryBin{reqs: &r}
	man.save[binID] = &b
	return &b
}

// Bin returns bin which binID pointing
func (man *MemoryBinManager) Bin(binID string) *MemoryBin {
	b, ok := man.save[binID]
	if !ok {
		return nil
	}
	return b
}
