package functools

import (
	"reflect"
	"errors"
)

/*
'Max' function returns the largest item in an iterable collection.
It compares items by 'Cmp' function. Function allows slices and arrays.

	Max(slice) interface{}
	MaxSafe(slice) (interface{}, err)
*/

func max(slice interface{}) interface{} {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Max")
	}

	if rv.Len() == 0 {
		raise(errors.New("The passed collection is an empty sequence"), "Max")
	}

	fn := reflect.ValueOf(interface{}(Cmp))
	// index: 0 - biggest value, 1 - current value
	params := [2]reflect.Value{rv.Index(0), rv.Index(0)}

	for i := 0; i < rv.Len(); i++ {
		params[1] = rv.Index(i)
		if fn.Call(params[:])[0].Int() < 0 {
			params[0] = rv.Index(i)
		}
	}

	return params[0].Interface()
}

func Max(slice interface{}) interface{} {
	return max(slice)
}

func MaxSafe(slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = max(slice)
	return
}
