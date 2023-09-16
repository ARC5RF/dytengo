package dytengo

import (
	"reflect"

	"github.com/d5/tengo/v2"
)

// DyArrayIterator is an iterator for an array.
type DyStructIterator struct {
	tengo.ObjectImpl
	v reflect.Value
	i int
	k int
	l int
}

// TypeName returns the name of the type.
func (i *DyStructIterator) TypeName() string {
	return "dy-struct-iterator"
}

func (i *DyStructIterator) String() string {
	return "<dy-struct-iterator>"
}

// IsFalsy returns true if the value of the type is falsy.
func (i *DyStructIterator) IsFalsy() bool {
	return true
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (i *DyStructIterator) Equals(tengo.Object) bool {
	return false
}

// Copy returns a copy of the type.
func (i *DyStructIterator) Copy() tengo.Object {
	return &DyStructIterator{v: i.v, i: i.i, l: i.l}
}

// Next returns true if there are more elements to iterate.
func (i *DyStructIterator) Next() bool {
	i.i++
	t := i.v.Type()
	for i.i < i.l {
		ft := t.Field(i.i - 1)
		if !ft.IsExported() {
			i.i++
			continue
		}
		break
	}
	return i.i <= i.l
}

// Key returns the key or index value of the current element.
func (i *DyStructIterator) Key() tengo.Object {
	t := i.v.Type()
	f := t.Field(i.i - 1)
	return &tengo.String{Value: f.Name}
}

// Value returns the value of the current element.
func (i *DyStructIterator) Value() tengo.Object {
	o, err := FromInterface(i.v.Field(i.i - 1).Interface())
	if err != nil {
		return tengo.UndefinedValue
	}
	return o
}
