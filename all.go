package functools

import (
	"reflect"
	"errors"
)

/*
'All' function returns 'true' if all elements of iterable 'slice' are 'true'
or if the slice is empty. Each element of the slice is converted to bool
by 'functools.ToBool' function. Function allowes slices and arrays.

	All(slice) bool
	AllSafe(slice) (bool, err)

'AllFunc' function applies 'func' parameter for each element of the iterable 'slice'
and returns 'true' if all results of 'func' calling are 'true' or if the slice is empty.
Function allowes slices and arrays.

	AllFunc(func, slice) bool
	AllFuncSafe(func, slice) (bool, err)
*/

func all(function, slice interface{}) bool {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "All")
	}

	fn := reflect.ValueOf(function)
	t := rv.Type().Elem()

	if !verifyAllFuncType(fn, t) {
		raise(errors.New("Function must be of type func(" + t.String() +
			") bool or func(interface{}) bool"), "All")
	}

	var param [1]reflect.Value

	for i := 0; i < rv.Len(); i++ {
		param[0] = rv.Index(i)

		if !fn.Call(param[:])[0].Bool() {
			return false
		}
	}

	return true
}

func verifyAllFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}

	return ((fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType) &&
		fn.Type().Out(0).Kind() == reflect.Bool)
}

func All(slice interface{}) bool {
	return all(ToBool, slice)
}

func AllSafe(slice interface{}) (result bool, err error) {
	defer except(&err)
	result = all(ToBool, slice)
	return
}

func AllFunc(function, slice interface{}) bool {
	return all(function, slice)
}

func AllFuncSafe(function, slice interface{}) (result bool, err error) {
	defer except(&err)
	result = all(function, slice)
	return
}
