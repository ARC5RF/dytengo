package dytengo

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/token"
)

type DyStruct struct {
	tengo.ObjectImpl
	Value reflect.Value
	t     reflect.Type
}

// TypeName returns the name of the type.
func (o *DyStruct) TypeName() string {
	return "dy-struct"
}

func (o *DyStruct) String() string {
	return "<dy-struct>"
}

// BinaryOp returns another object that is the result of a given binary
// operator and a right-hand side object.
func (o *DyStruct) BinaryOp(_ token.Token, _ tengo.Object) (tengo.Object, error) {
	return nil, tengo.ErrInvalidOperator
}

// Copy returns a copy of the type.
func (o *DyStruct) Copy() tengo.Object {
	return nil
}

// IsFalsy returns true if the value of the type is falsy.
func (o *DyStruct) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type is equal to the value of
// another object.
func (o *DyStruct) Equals(x tengo.Object) bool {
	return reflect.DeepEqual(o, x)
}

var ErrIndexNotAssignable = errors.New("index cannot be set")

// IndexGet returns an element at a given index.
func (o *DyStruct) IndexGet(key tengo.Object) (res tengo.Object, err error) {
	_str, ok := key.(*tengo.String)
	if ok {
		str := _str.Value
		fmt.Println(str)
		v := o.Value
		if o.Value.Kind() == reflect.Pointer {
			v = reflect.Indirect(v)
		}
		fv := v.FieldByName(str)
		if fv != (reflect.Value{}) {
			return FromInterface(fv.Interface())
		}

		mv := o.Value.MethodByName(str)
		fmt.Println("mv", mv.Type().String())
		if mv != (reflect.Value{}) {
			return FromInterface(mv.Interface())
		}
		return tengo.UndefinedValue, nil
	}

	return nil, tengo.ErrInvalidIndexType
}

// IndexSet sets an element at a given index.
func (o *DyStruct) IndexSet(key, val tengo.Object) error {
	_str, ok := key.(*tengo.String)
	if ok {
		str := _str.Value
		v := o.Value
		if o.Value.Kind() == reflect.Pointer {
			v = reflect.Indirect(v)
		}
		fv := v.FieldByName(str)
		ft := fv.Type()
		if !fv.CanSet() {
			return ErrIndexNotAssignable
		}
		iv := reflect.ValueOf(ToInterface(val))
		it := iv.Type()

		if it.String() != ft.String() {
			if !iv.CanConvert(ft) {
				return tengo.ErrInvalidIndexValueType
			}
			iv = iv.Convert(ft)
		}
		fv.Set(iv)

		return nil
	}

	return tengo.ErrInvalidIndexType
}

// Iterate returns an iterator.
func (o *DyStruct) Iterate() tengo.Iterator {
	v := o.Value
	if o.Value.Kind() == reflect.Pointer {
		v = reflect.Indirect(v)
	}
	n := v.NumField()
	return &DyStructIterator{
		v: v,
		l: n,
	}
}

// CanIterate returns whether the Object can be Iterated.
func (o *DyStruct) CanIterate() bool {
	if o.Value.Kind() == reflect.Pointer {
		return !o.Value.IsNil()
	}
	return true
}
