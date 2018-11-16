package functools

import (
	"reflect"
	"errors"
)

/*
'Avg' function returns average value of all items in 'slice' as 'float64'.
The items should be numbers. Function allows slices and arrays.

	Avg(slice) float64
	AvgSafe(slice) (float64, err)
*/

func avg(slice interface{}) float64 {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		raise(errors.New("The passed collection is not a slice or an array"), "Avg")
	}

	if rv.Len() == 0 {
		raise(errors.New("The passed collection is an empty sequence"), "Avg")
	}

	fn := reflect.ValueOf(interface{}(divideAB))
	var params [2]reflect.Value
	params[0] = reflect.ValueOf(Sum(slice))
	params[1] = reflect.ValueOf(interface{}(rv.Len()))

	return fn.Call(params[:])[0].Float()
}

func Avg(slice interface{}) float64 {
	return avg(slice)
}

func AvgSafe(slice interface{}) (result float64, err error) {
	defer except(&err)
	result = avg(slice)
	return
}

func divideAB(a, b interface{}) float64 {
	if a == nil || b == nil {
		raise(errors.New("You cannot divide nil values"), "Avg")
	}

	rvA, rvB := reflect.ValueOf(a), reflect.ValueOf(b)

	var casted_a, casted_b float64

	switch a.(type) {
		case int, int8, int16, int32, int64:
			casted_a = float64(rvA.Int())
		case uint, uint8, uint16, uint32, uint64:
			casted_a = float64(rvA.Uint())
		case float32, float64:
			casted_a = rvA.Float()
		default:
			k := rvA.Type().Kind()
			raise(errors.New("Unexpected type (" + k.String() + ") of first value"), "Avg")
	}

	switch b.(type) {
		case int, int8, int16, int32, int64:
			casted_b = float64(rvB.Int())
		case uint, uint8, uint16, uint32, uint64:
			casted_b = float64(rvB.Uint())
		case float32, float64:
			casted_b = rvB.Float()
		default:
			k := rvB.Type().Kind()
			raise(errors.New("Unexpected type (" + k.String() + ") of second value"), "Avg")
	}

	return casted_a / casted_b
}
