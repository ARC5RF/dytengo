package dytengo

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/token"
)

type DyMap struct {
	tengo.ObjectImpl
	Value reflect.Value
	t     reflect.Type
}

func (o *DyMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value.Interface())
}

func (o *DyMap) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, o.Value.Interface())
}

// TypeName returns the name of the type.
func (o *DyMap) TypeName() string {
	return "dy-map"
}

func (o *DyMap) String() string {
	return "<dy-map>"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *DyMap) BinaryOp(_ token.Token, _ tengo.Object) (tengo.Object, error) {
	return nil, tengo.ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *DyMap) Copy() tengo.Object {
	obj, err := FromInterface(o.Value.Interface())
	if err != nil {
		panic(err)
	}
	return obj
}

// IsFalsy returns true if the value of the type is falsy.
func (o *DyMap) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *DyMap) Equals(x tengo.Object) bool {
	return reflect.DeepEqual(o, x)
}

// IndexGet returns an element at a given index.
func (o *DyMap) IndexGet(index tengo.Object) (res tengo.Object, err error) {
	if o.Value.IsNil() {
		return tengo.UndefinedValue, nil
	}

	ii := ToInterface(index)
	iv := reflect.ValueOf(ii)

	it := iv.Type()

	kt := o.t.Key()

	if kt.String() != it.String() {
		if !iv.CanConvert(kt) {
			return nil, tengo.ErrInvalidIndexType
		}
		iv = iv.Convert(kt)
	}

	v := o.Value.MapIndex(iv)
	if v != (reflect.Value{}) {
		return FromInterface(v.Interface())
	}

	return tengo.UndefinedValue, nil
}

var ErrMapIsNil = errors.New("cannot index nil map")

// IndexSet sets an element at a given index.
func (o *DyMap) IndexSet(index, val tengo.Object) (err error) {
	if o.Value.IsNil() {
		return ErrMapIsNil
	}

	ii := ToInterface(index)
	iv := reflect.ValueOf(ii)

	it := iv.Type()
	kt := o.t.Key()

	if kt.String() != it.String() {
		if !iv.CanConvert(kt) {
			return tengo.ErrInvalidIndexType
		}
		iv = iv.Convert(kt)
	}

	vt := o.t.Elem()
	ei := ToInterface(val)
	ev := reflect.ValueOf(ei)
	et := ev.Type()

	if vt.String() != et.String() {
		if !ev.CanConvert(vt) {
			return tengo.ErrInvalidIndexValueType
		}
		ev = ev.Convert(vt)
	}

	v := o.Value.MapIndex(iv)
	if v != (reflect.Value{}) {
		o.Value.SetMapIndex(iv, ev)
		return nil
	}

	return tengo.ErrNotIndexAssignable
}

// Iterate returns an iterator.
func (o *DyMap) Iterate() tengo.Iterator {
	return &DyMapIterator{v: o.Value.MapRange()}
}

// CanIterate returns whether the Object can be Iterated.
func (o *DyMap) CanIterate() bool {
	return !o.Value.IsNil()
}
