package functools

import (
	"reflect"
	"errors"
)

/*
'Filter' function construct a slice from those elements of iterable 'slice'
which can be represented as bool 'true'. Each element of the slice is converted to bool
by 'functools.ToBool' function. Function allowes slices and arrays.

	Filter(slice) slice
	FilterSafe(slice) (slice, err)

'FilterFunc' function construct a slice from those elements of iterable 'slice'
for which 'func' returns 'true'. Function allowes slices and arrays.

	FilterFunc(func, slice) slice
	FilterFuncSafe(func, slice) (slice, err)
*/

func filter(function, slice interface{}) interface{} {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Filter")
	}

	fn := reflect.ValueOf(function)
	t := rv.Type().Elem()

	if !verifyFilterFuncType(fn, t) {
		raise(errors.New("Function must be of type func(" + t.String() +
			") bool or func(interface{}) bool"), "Filter")
	}

	var param [1]reflect.Value
	out := reflect.MakeSlice(rv.Type(), 0, rv.Len())

	for i := 0; i < rv.Len(); i++ {
		param[0] = rv.Index(i)

		if fn.Call(param[:])[0].Bool() {
			out = reflect.Append(out, rv.Index(i))
		}
	}

	return out.Interface()
}

func verifyFilterFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}

	return ((fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType) &&
		fn.Type().Out(0).Kind() == reflect.Bool)
}

func Filter(slice interface{}) interface{} {
	return filter(ToBool, slice)
}

func FilterSafe(slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = filter(ToBool, slice)
	return
}

func FilterFunc(function, slice interface{}) interface{} {
	return filter(function, slice)
}

func FilterFuncSafe(function, slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = filter(function, slice)
	return
}
