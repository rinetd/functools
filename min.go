package functools

import (
	"reflect"
	"errors"
)

/*
'Min' function returns the smallest item in an iterable collection.
It compares items by 'Cmp' function. Function allows slices and arrays.

	Min(slice) interface{}
	MinSafe(slice) (interface{}, err)
*/

func min(slice interface{}) interface{} {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Max")
	}

	if rv.Len() == 0 {
		raise(errors.New("The passed collection is an empty sequence"), "Max")
	}

	fn := reflect.ValueOf(interface{}(Cmp))
	// index: 0 - least value, 1 - current value
	params := [2]reflect.Value{rv.Index(0), rv.Index(0)}

	for i := 0; i < rv.Len(); i++ {
		params[1] = rv.Index(i)
		if fn.Call(params[:])[0].Int() > 0 {
			params[0] = rv.Index(i)
		}
	}

	return params[0].Interface()
}

func Min(slice interface{}) interface{} {
	return min(slice)
}

func MinSafe(slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = min(slice)
	return
}
