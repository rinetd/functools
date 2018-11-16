package functools

import (
	"reflect"
	"errors"
)

/*
'Any' function returns 'true' if any element of the iterable 'slice' is 'true'.
If the slice is empty, returns 'false'. Each element of the slice is converted to bool
by 'functools.ToBool' function. Function allows slices and arrays.

	Any(slice) bool
	AnySafe(slice) (bool, err)

'AnyFunc' function applies 'func' parameter for each element of the iterable 'slice'
and returns 'true' if any result of 'func' calling is 'true'. If the slice is empty,
returns 'false'. Function allows slices and arrays.

	AnyFunc(func, slice) bool
	AnyFuncSafe(func, slice) (bool, err)
*/

func any(function, slice interface{}) bool {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Any")
	}

	fn := reflect.ValueOf(function)
	t := rv.Type().Elem()

	if !verifyAnyFuncType(fn, t) {
		raise(errors.New("Function must be of type func(" + t.String() +
			") bool or func(interface{}) bool"), "Any")
	}

	var param [1]reflect.Value

	for i := 0; i < rv.Len(); i++ {
		param[0] = rv.Index(i)

		if fn.Call(param[:])[0].Bool() {
			return true
		}
	}

	return false
}

func verifyAnyFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}

	return ((fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType) &&
		fn.Type().Out(0).Kind() == reflect.Bool)
}

func Any(slice interface{}) bool {
	return any(ToBool, slice)
}

func AnySafe(slice interface{}) (result bool, err error) {
	defer except(&err)
	result = any(ToBool, slice)
	return
}

func AnyFunc(function, slice interface{}) bool {
	return any(function, slice)
}

func AnyFuncSafe(function, slice interface{}) (result bool, err error) {
	defer except(&err)
	result = any(function, slice)
	return
}
