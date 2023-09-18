package dytengo

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/token"
)

type DyArray struct {
	tengo.ObjectImpl
	Value reflect.Value
	t     reflect.Type
}

func (o *DyArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value.Interface())
}

func (o *DyArray) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, o.Value.Interface())
}

// TypeName returns the name of the type.
func (o *DyArray) TypeName() string {
	return "dy-array"
}

func (o *DyArray) String() string {
	return "<dy-array>"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *DyArray) BinaryOp(_ token.Token, _ tengo.Object) (tengo.Object, error) {
	return nil, tengo.ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *DyArray) Copy() tengo.Object {
	obj, err := FromInterface(o.Value.Interface())
	if err != nil {
		panic(err)
	}
	return obj
}

// IsFalsy returns true if the value of the type is falsy.
func (o *DyArray) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *DyArray) Equals(x tengo.Object) bool {
	return reflect.DeepEqual(o, x)
}

var ErrArrayNil = errors.New("cannot index nil array")

// IndexGet returns an element at a given index.
func (o *DyArray) IndexGet(index tengo.Object) (res tengo.Object, err error) {
	if o.Value.IsNil() {
		return nil, ErrArrayNil
	}

	ii := ToInterface(index)
	iv := reflect.ValueOf(ii)
	it := iv.Type()

	var k int
	kv := reflect.ValueOf(k)
	kt := kv.Type()

	if kt.String() != it.String() {
		if !iv.CanConvert(kt) {
			return nil, tengo.ErrInvalidIndexType
		}
		iv = iv.Convert(kt)
	}
	k = iv.Interface().(int)

	if k < 0 || k >= o.Value.Len() {
		return nil, tengo.ErrIndexOutOfBounds
	}

	e := o.Value.Index(k)

	return FromInterface(e.Interface())
}

// IndexSet sets an element at a given index.
func (o *DyArray) IndexSet(index, val tengo.Object) (err error) {
	if o.Value.IsNil() {
		return ErrArrayNil
	}

	ii := ToInterface(index)
	iv := reflect.ValueOf(ii)
	it := iv.Type()

	ei := ToInterface(val)
	ev := reflect.ValueOf(ei)

	et := ev.Type()
	vt := o.t.Elem()

	if vt.String() != et.String() {
		if !ev.CanConvert(vt) {
			return tengo.ErrInvalidIndexValueType
		}
		ev = ev.Convert(vt)
	}

	var k int
	kv := reflect.ValueOf(k)
	kt := kv.Type()

	if kt.String() != it.String() {
		if !iv.CanConvert(kt) {
			return tengo.ErrInvalidIndexType
		}
		iv = iv.Convert(kt)
	}
	k = iv.Interface().(int)

	if k < 0 || k >= o.Value.Len() {
		return tengo.ErrIndexOutOfBounds
	}

	e := o.Value.Index(k)
	e.Set(ev)

	return nil
}

// Iterate returns an iterator.
func (o *DyArray) Iterate() tengo.Iterator {
	return &DyArrayIterator{
		v: o.Value,
		l: o.Value.Len(),
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *DyArray) CanIterate() bool {
	return !o.Value.IsNil()
}
