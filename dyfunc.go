package dytengo

import (
	"reflect"

	"github.com/d5/tengo/v2"
)

type DyFunc struct {
	tengo.ObjectImpl
	Value reflect.Value
	t     reflect.Type
}

// TypeName returns the name of the type.
func (o *DyFunc) TypeName() string {
	return "dy-func"
}

func (o *DyFunc) String() string {
	return "<dy-func>"
}

func ins(args ...tengo.Object) []reflect.Value {
	ins := []reflect.Value{}

	for _, arg := range args {
		a := ToInterface(arg)

		v := reflect.ValueOf(a)
		ins = append(ins, v)
	}
	return ins
}

// Call takes an arbitrary number of arguments and returns a return value
// and/or an error.
func (o *DyFunc) Call(args ...tengo.Object) (ret tengo.Object, err error) {
	ft := o.Value.Type()
	oc := ft.NumOut()

	outs := o.Value.Call(ins(args...))
	if oc == 0 {
		return tengo.UndefinedValue, nil
	}
	if oc == 1 {
		return FromInterface(outs[0].Interface())
	}
	output := &tengo.Array{Value: []tengo.Object{}}
	for _, o := range outs {
		o, err := FromInterface(o.Interface())
		if err != nil {
			return nil, err
		}
		output.Value = append(output.Value, o)
	}
	return output, nil
}

// CanCall returns whether the Object can be Called.
func (o *DyFunc) CanCall() bool {
	return true
}
