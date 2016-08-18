// Practice: indexing

package main

import "log"

// Lenable requeres implementation
type Lenable interface {
	At(int) interface{}
	Len() int
}

// LogLenableLen works on Lenable
func LogLenableLen(l Lenable) {
	log.Println("#LogLenableLen", l.Len())
}

// LogLenable0 works on Lenable
func LogLenable0(l Lenable) {
	// invalid operation: l[0] (type Lenable does not support indexing)
	// log.Println("#LogLenable0", l[0])

	log.Println("#LogLenable0", l.At(0))
}

// Record represents example data
type Record struct {
	ID int
}

// Records represents collection
type Records []Record

// At implements Lenable interface
func (rr *Records) At(at int) interface{} {
	if len(*rr) > at {
		return (*rr)[at]
	}
	return nil
}

// Len implements Lenable interface
func (rr *Records) Len() int {
	return len(*rr)
}

func main() {
	rr := Records{{1}, {2}}
	log.Println("#rr", rr, rr.Len())
	log.Println("#rr[0]", rr[0])

	LogLenableLen(&rr)
	LogLenable0(&rr)
}
