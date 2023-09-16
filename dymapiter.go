package dytengo

import (
	"reflect"

	"github.com/d5/tengo/v2"
)

// DyMapIterator represents an iterator for the map.
type DyMapIterator struct {
	tengo.ObjectImpl
	v *reflect.MapIter
}

// TypeName returns the name of the type.
func (i *DyMapIterator) TypeName() string {
	return "dy-map-iterator"
}

func (i *DyMapIterator) String() string {
	return "<dy-map-iterator>"
}

// IsFalsy returns true if the value of the type is falsy.
func (i *DyMapIterator) IsFalsy() bool {
	return true
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (i *DyMapIterator) Equals(tengo.Object) bool {
	return false
}

// Copy returns a copy of the type.
func (i *DyMapIterator) Copy() tengo.Object {
	return &DyMapIterator{v: i.v}
}

// Next returns true if there are more elements to iterate.
func (i *DyMapIterator) Next() bool {
	return i.v.Next()
}

// Key returns the key or index value of the current element.
func (i *DyMapIterator) Key() tengo.Object {
	o, err := FromInterface(i.v.Key().Interface())
	if err != nil {
		return tengo.UndefinedValue
	}
	return o
}

// Value returns the value of the current element.
func (i *DyMapIterator) Value() tengo.Object {
	o, err := FromInterface(i.v.Value().Interface())
	if err != nil {
		return tengo.UndefinedValue
	}
	return o
}
