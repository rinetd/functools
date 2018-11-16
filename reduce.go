package functools

import (
	"reflect"
	"errors"
)

/*
'Reduce' applies 'func' of two arguments cumulatively to the items of iterable 'slice',
from left to right, so as to reduce the iterable to a single 'acc' value.
Function allowes slices and arrays.

	Reduce(func, slice, acc) acc
	ReduceSafe(func, slice, acc) (acc, err)
 */

func reduce(function, slice, acc interface{}) interface{} {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Reduce")
	}

	fn := reflect.ValueOf(function)
	t := rv.Type().Elem()

	if !verifyReduceFuncType(fn, t) {
		raise(errors.New("Function must be of type func(" + t.String() + ", " + t.String() +
			") " + t.String() + " or func(interface{}, interface{}) " + t.String()), "Reduce")
	}

	if t != reflect.TypeOf(acc) {
		raise(errors.New("The type of accumulator and the type of elements of slice should be the same"), "Reduce")
	}

	var param [2]reflect.Value
	out := reflect.ValueOf(acc)

	for i := 0; i < rv.Len(); i++ {
		param[0] = out
		param[1] = rv.Index(i)
		out = fn.Call(param[:])[0]
	}

	return out.Interface()
}

func verifyReduceFuncType(fn reflect.Value, elemType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 2 || fn.Type().NumOut() != 1 {
		return false
	}

	if fn.Type().In(0) != fn.Type().In(1) || fn.Type().Out(0) != elemType {
		return false
	}

	return fn.Type().In(0).Kind() == reflect.Interface || elemType == fn.Type().In(0)
}

func Reduce(function, slice, acc interface{}) interface{} {
	return reduce(function, slice, acc)
}

func ReduceSafe(function, slice, acc interface{}) (result interface{}, err error) {
	defer except(&err)
	result = reduce(function, slice, acc)
	return
}
