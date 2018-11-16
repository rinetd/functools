package functools

import (
	"errors"
	"reflect"
)

/*
'Map' applies 'func' to every item of iterable 'slice' and return a slice of the results.
Function allowes slices and arrays.

	Map(func, slice) slice
	MapSafe(func, slice) (slice, err)
*/
func Apply(function, slice interface{}) interface{} {
	return _map(function, slice)
}
func ApplySafe(function, slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = _map(function, slice)
	return
}
func Map(function, slice interface{}) interface{} {
	return _map(function, slice)
}

func MapSafe(function, slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = _map(function, slice)
	return
}

func _map(function, slice interface{}) interface{} { // Resolved conflict with built-in map statement
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Map")
	}

	fn := reflect.ValueOf(function)
	t := rv.Type().Elem()

	if !verifyMapFuncType(fn, t) {
		raise(errors.New("Function must be of type func("+t.String()+
			") interface{} or func(interface{}) interface{}"), "Map")
	}

	var param [1]reflect.Value
	out := reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), 0, rv.Len())

	for i := 0; i < rv.Len(); i++ {
		param[0] = rv.Index(i)
		out = reflect.Append(out, fn.Call(param[:])[0])
	}

	return out.Interface()
}

func verifyMapFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}

	return fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType
}
