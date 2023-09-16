package dytengo

import (
	"reflect"

	"github.com/d5/tengo/v2"
)

// DyArrayIterator is an iterator for an array.
type DyArrayIterator struct {
	tengo.ObjectImpl
	v reflect.Value
	i int
	l int
}

// TypeName returns the name of the type.
func (i *DyArrayIterator) TypeName() string {
	return "dy-array-iterator"
}

func (i *DyArrayIterator) String() string {
	return "<dy-array-iterator>"
}

// IsFalsy returns true if the value of the type is falsy.
func (i *DyArrayIterator) IsFalsy() bool {
	return true
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (i *DyArrayIterator) Equals(tengo.Object) bool {
	return false
}

// Copy returns a copy of the type.
func (i *DyArrayIterator) Copy() tengo.Object {
	return &DyArrayIterator{v: i.v, i: i.i, l: i.l}
}

// Next returns true if there are more elements to iterate.
func (i *DyArrayIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

// Key returns the key or index value of the current element.
func (i *DyArrayIterator) Key() tengo.Object {
	return &tengo.Int{Value: int64(i.i - 1)}
}

// Value returns the value of the current element.
func (i *DyArrayIterator) Value() tengo.Object {
	o, err := FromInterface(i.v.Index(i.i - 1).Interface())
	if err != nil {
		return tengo.UndefinedValue
	}
	return o
}
