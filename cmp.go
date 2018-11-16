package functools

import (
	"reflect"
	"errors"
	"unicode/utf8"
)

/*
'Cmp' function compares two objects 'a' and 'b' and returns an 'int8' according
to outcome. The return value is negative if 'a' < 'b', zero if 'a' == 'b'
and strictly positive if 'a' > 'b'. For strings and collection function compares
their lengths. For bool values 'true' > 'false'.

	Cmp(a, b) int8
	CmpSafe(a, b) (int8, err)
*/

func cmp(a, b interface{}) int8 {
	if a == nil || b == nil {
		if a == nil && b == nil {
			return 0
		} else if a == nil && b != nil {
			return -1
		} else {
			return 1
		}
	}

	rvA, rvB := reflect.ValueOf(a), reflect.ValueOf(b)

	if rvA.Type() != rvB.Type() {
		raise(errors.New("Both parameters should have the same types"), "Cmp")
	}

	var casted_a, casted_b float64

	switch a.(type) {
		case int, int8, int16, int32, int64:
			casted_a, casted_b = float64(rvA.Int()), float64(rvB.Int())
		case uint, uint8, uint16, uint32, uint64:
			casted_a, casted_b = float64(rvA.Uint()), float64(rvB.Uint())
		case float32, float64:
			casted_a, casted_b = rvA.Float(), rvB.Float()
		case string:
			casted_a = float64(utf8.RuneCountInString(rvA.String()))
			casted_b = float64(utf8.RuneCountInString(rvB.String()))
		case bool:
			if rvA.Bool() {
				casted_a = 1.0
			}
			if rvB.Bool() {
				casted_b = 1.0
			}
		default:
			k := rvA.Type().Kind()
			if k == reflect.Array || k == reflect.Slice || k == reflect.Map {
				casted_a, casted_b = float64(rvA.Len()), float64(rvB.Len())
			} else {
				raise(errors.New("Unexpected type (" + k.String() + ") of values"), "Cmp")
			}
	}

	if casted_a == casted_b {
		return 0
	} else if casted_a < casted_b {
		return -1
	} else {
		return 1
	}
}

func Cmp(a, b interface{}) int8 {
	return cmp(a, b)
}

func CmpSafe(a, b interface{}) (result int8, err error) {
	defer except(&err)
	result = cmp(a, b)
	return
}
