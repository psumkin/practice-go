// Practice: indexing

package main

import "log"

// Indexed requeres implementation
type Indexed interface {
	At(int) interface{}
	Len() int
}

// LogIndexed works on Indexed
func LogIndexed(l Indexed) {
	// invalid argument l (type Indexed) for len
	// log.Println("#IndexedLen", len(l))

	log.Println("#IndexedLen", l.Len())

	// invalid operation: l[0] (type Indexed does not support indexing)
	// log.Println("#IndexedAt0", l[0])

	log.Println("#IndexedAt0", l.At(0))
}

// T represents example data
type T struct {
	i int
}

// Map represents collection
type Map map[int]T

// At implements Indexed interface
func (tt Map) At(at int) interface{} {
	if len(tt) > at {
		return tt[at]
	}
	return nil
}

// Len implements Indexed interface
func (tt Map) Len() int {
	return len(tt)
}

// PMap represents collection
type PMap map[int]*T

// At implements Indexed interface
func (tt PMap) At(at int) interface{} {
	if len(tt) > at {
		return tt[at]
	}
	return nil
}

// Len implements Indexed interface
func (tt PMap) Len() int {
	return len(tt)
}

// Slice represents collection
type Slice []T

// At implements Indexed interface
func (tt Slice) At(at int) interface{} {
	if len(tt) > at {
		return tt[at]
	}
	return nil
}

// Len implements Indexed interface
func (tt Slice) Len() int {
	return len(tt)
}

// PSlice represents collection
type PSlice []*T

// At implements Indexed interface
func (tt PSlice) At(at int) interface{} {
	if len(tt) > at {
		return tt[at]
	}
	return nil
}

// Len implements Indexed interface
func (tt PSlice) Len() int {
	return len(tt)
}

func main() {
	map1 := Map{0: {11}, 1: {22}}
	log.Println("#map1", map1, len(map1), map1[0])

	LogIndexed(map1)
	LogIndexed(&map1)

	pmap1 := PMap{0: {11}, 1: {22}}
	log.Println("#pmap1", pmap1, len(pmap1), pmap1[0])

	LogIndexed(pmap1)
	LogIndexed(&pmap1)

	slice1 := Slice{{11}, {22}}
	log.Println("#slice1", slice1, len(slice1), slice1[0])

	LogIndexed(slice1)
	LogIndexed(&slice1)

	pslice1 := PSlice{{11}, {22}}
	log.Println("#pslice1", pslice1, len(pslice1), pslice1[0])

	LogIndexed(pslice1)
	LogIndexed(&pslice1)

	var pmap2 PMap
	log.Println("#pmap2", pmap2, len(pmap2), pmap2[0], pmap2 == nil)
	// panic: assignment to entry in nil map
	// pmap2[0] = &T{11}

	pmap3 := make(PMap)
	log.Println("#pmap3", pmap3, len(pmap3), pmap3[0], pmap3 == nil)
	pmap3[0] = &T{11}
	LogIndexed(pmap3)

}
