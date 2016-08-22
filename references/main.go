// Practice: references in interfaces

package main

import (
	"log"

	"github.com/satori/go.uuid"
)

// Identifiable requeres implementation
type Identifiable interface {
	GetID() uuid.UUID
	SetID(uuid.UUID)
	TouchID(uuid.UUID)
}

// GetID works on Identifiable
func GetID(i Identifiable) {
	log.Println("#GetID", i.GetID())
}

// SetID works on Identifiable
func SetID(i Identifiable) {
	i.SetID(uuid.NewV4())
}

// TouchID works on Identifiable
func TouchID(i Identifiable) {
	i.TouchID(uuid.NewV4())
}

// Update works on *Identifiable
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

// GetID implements Identifiable interface on pointer
func (r *Record) GetID() uuid.UUID {
	return r.ID
}

// SetID implements Identifiable interface on pointer
func (r *Record) SetID(id uuid.UUID) {
	r.ID = id
}

// TouchID implements Identifiable interface on struct
func (r Record) TouchID(id uuid.UUID) {
	r.ID = id
	log.Println("#TouchID", r.ID)
}

// Records represents collection
type Records []Record

// Check types at compiling
var _ Identifiable = &Record{}

func main() {
	rec1 := Record{uuid.NewV4()}
	rec2 := NewRecord()
	log.Println("#list #1", rec1, rec2)

	GetID(&rec1)
	GetID(rec2)

	SetID(&rec1)
	SetID(rec2)
	log.Println("#list #2", rec1, rec2)

	TouchID(&rec1)
	TouchID(rec2)
	log.Println("#list #3", rec1, rec2)

	var i1, i2 Identifiable
	i1 = &rec1
	i2 = rec2
	Update(&i1)
	Update(&i2)

	rr := Records{rec1, *rec2}
	log.Println("#rr", rr)
}
