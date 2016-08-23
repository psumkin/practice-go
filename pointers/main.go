// Practice: references in interfaces

package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

// Identifiable requeres implementation
type Identifiable interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
	TouchID(uuid.UUID)
}

// Identify works on Identifiable
func Identify(i Identifiable) {
	fmt.Println("#Identify", i.GetID())
	i.TouchID(uuid.NewV4())
	fmt.Println("#", i.GetID())
	i.SetID(uuid.NewV4())
	fmt.Println("#", i.GetID())
	fmt.Println()
}

// Update works on *Identifiable
// https://golang.org/doc/faq#pointer_to_interface
func Update(i *Identifiable) {
	(*i).SetID(uuid.NewV4())
}

// Record represents example data
type Record struct {
	ID uuid.UUID
}

// NewRecord implements constructor
// inspired by http://stackoverflow.com/a/18125682/4825998
func NewRecord() *Record {
	// return &Record{uuid.NewV4()}
	p := new(Record)
	p.ID = uuid.NewV4()
	return p
}

// GetID implements Identifiable interface on struct
func (r Record) GetID() uuid.UUID {
	return r.ID
}

// SetID implements Identifiable interface on pointer
func (r *Record) SetID(id uuid.UUID) {
	r.ID = id
}

// TouchID implements Identifiable interface on struct
// https://golang.org/doc/faq#methods_on_values_or_pointers
func (r Record) TouchID(id uuid.UUID) {
	r.ID = id
	fmt.Println("#TouchID", r.ID)
}

// Records represents collection
type Records []Record

// Check types at compiling
var _ Identifiable = &Record{}

func main() {
	rec1 := Record{uuid.NewV4()}
	rec2 := NewRecord()
	fmt.Println("#list #1", rec1, rec2)

	Identify(&rec1)
	Identify(rec2)

	fmt.Println("#list #2", rec1, rec2)

	(*Record).SetID(&rec1, uuid.NewV4())
	Identifiable.SetID(rec2, uuid.NewV4())
	fmt.Println("#list #4", rec1, rec2)

	var i1, i2 Identifiable
	i1 = &rec1
	i2 = rec2
	Update(&i1)
	Update(&i2)

	rr := Records{rec1, *rec2}
	fmt.Println("#rr", rr)
}
