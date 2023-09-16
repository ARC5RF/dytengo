package dytengo

import (
	"errors"
	"reflect"

	"github.com/d5/tengo/v2"
)

type DyScript struct {
	*tengo.Script
}

func (s *DyScript) Add(name string, value any) error {
	o, err := FromInterface(value)
	if err != nil {
		return err
	}
	return s.Script.Add(name, o)
}

func NewScript(src []byte) *DyScript {
	self := &DyScript{tengo.NewScript(src)}
	return self
}

var ErrStructNotPointer = errors.New("from interface recieved non pointer struct")

func FromInterface(v any) (tengo.Object, error) {
	o, err := tengo.FromInterface(v)
	if err != nil {
		rv := reflect.ValueOf(v)
		rk := rv.Kind()

		uv := rv
		if rk == reflect.Interface || rk == reflect.Pointer {
			uv = reflect.Indirect(rv)
			rk = uv.Kind()
		}
		switch rk {
		case reflect.Array, reflect.Slice:
			return &DyArray{Value: rv, t: uv.Type()}, nil
		case reflect.Map:
			return &DyMap{Value: rv, t: uv.Type()}, nil
		case reflect.Struct:
			return &DyStruct{Value: rv, t: uv.Type()}, nil
		case reflect.Func:
			return &DyFunc{Value: rv, t: uv.Type()}, nil
		}
	}
	return o, err
}

func ToInterface(o tengo.Object) any {
	res := tengo.ToInterface(o)
	switch v := res.(type) {
	case *DyArray:
		return v.Value.Interface()
	case *DyMap:
		return v.Value.Interface()
	case *DyStruct:
		return v.Value.Interface()
	case *DyFunc:
		return v.Value.Interface()
	}
	return res
}
