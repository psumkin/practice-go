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
}

// LogID works on Identifiable
func LogID(i Identifiable) {
	log.Println("#LogID", i.GetID())
}

// UpdateID works on Identifiable
func UpdateID(i Identifiable) {
	i.SetID(uuid.NewV4())
}

// Update works on *Identifiable - compiling error
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

// GetID implements Identifiable interface
func (r *Record) GetID() uuid.UUID {
	return r.ID
}

// SetID implements Identifiable interface
func (r *Record) SetID(id uuid.UUID) {
	r.ID = id
}

// Records represents collection
type Records []Record

// Getter requeres implementation
type Getter interface {
	Get() uuid.UUID
}

// Setter requeres implementation
type Setter interface {
	Set(uuid.UUID)
}

// LogGetter works on Getter
func LogGetter(i Getter) {
	log.Println("#LogGetter", i.Get())
}

// SetSetter works on Getter
func SetSetter(i Setter) {
	i.Set(uuid.NewV4())
}

// Set implements Setter interface
func (r *Record) Set(id uuid.UUID) {
	r.ID = id
}

// Get implements Getter interface
func (r Record) Get() uuid.UUID {
	return r.ID
}

// Check types at compiling
var _ Identifiable = &Record{}
var _ Getter = Record{}
var _ Setter = &Record{}

func main() {
	rec1 := Record{uuid.NewV4()}
	rec2 := NewRecord()
	log.Println("#list", rec1, rec2)

	LogID(&rec1)
	LogID(rec2)

	UpdateID(&rec1)
	UpdateID(rec2)
	LogID(&rec1)
	LogID(rec2)

	// Update(&rec1)
	// Update(rec2)
	// LogID(&rec1)
	// LogID(rec2)

	rr := Records{rec1, *rec2}
	log.Println("#rr", rr)

	LogGetter(rec1)
	LogGetter(rec2)
	SetSetter(&rec1)
	SetSetter(rec2)
	LogGetter(rec1)
	LogGetter(rec2)
}
