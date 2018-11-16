package functools

import (
	"reflect"
	"errors"
)

/*
'Sum' function sums items of `slice` from left to right and returns total.
The items should be numbers. Function allows slices and arrays.

	Sum(slice) interface{}
	SumSafe(slice) (interface{}, err)
*/

func sum(slice interface{}) interface{} {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Sum")
	}

	if rv.Len() == 0 {
		raise(errors.New("The passed collection is an empty sequence"), "Sum")
	}

	fn := reflect.ValueOf(interface{}(sumAB))
	var params [2]reflect.Value
	params[0] = rv.Index(0)

	for i := 1; i < rv.Len(); i++ {
		params[1] = rv.Index(i)
		params[0] = fn.Call(params[:])[0]
	}

	return params[0].Interface()
}

func Sum(slice interface{}) interface{} {
	return sum(slice)
}

func SumSafe(slice interface{}) (result interface{}, err error) {
	defer except(&err)
	result = sum(slice)
	return
}

// TODO: How can I append two numbers by reflect package?
func sumAB(a, b interface{}) interface{} {
	if a == nil || b == nil {
		raise(errors.New("You cannot append nil values"), "Sum")
	}

	rvA, rvB := reflect.ValueOf(a), reflect.ValueOf(b)

	if rvA.Type() != rvB.Type() {
		raise(errors.New("Both parameters should have the same types"), "Sum")
	}

	switch a.(type) {
		case int:
			return a.(int) + b.(int)
		case int8:
			return a.(int8) + b.(int8)
		case int16:
			return a.(int16) + b.(int16)
		case int32:
			return a.(int32) + b.(int32)
		case int64:
			return a.(int64) + b.(int64)
		case uint:
			return a.(int) + b.(int)
		case uint8:
			return a.(int8) + b.(int8)
		case uint16:
			return a.(int16) + b.(int16)
		case uint32:
			return a.(int32) + b.(int32)
		case uint64:
			return a.(int64) + b.(int64)
		case float32:
			return a.(float32) + b.(float32)
		case float64:
			return a.(float64) + b.(float64)
		default:
			k := rvA.Type().Kind()
			raise(errors.New("Unexpected type (" + k.String() + ") of values"), "Sum")
	}
	return 0.0
}
