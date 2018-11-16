package functools

import (
	"reflect"
	"errors"
)

/*
'ToBool' function returns 'true' if numeric 'value' parameter isn't equal to zero,
string or iterable collections aren't empty, and for bool 'value' parameter it returns
original value.

	ToBool(value) bool
	ToBoolSafe(value) (bool, err)
*/

func toBool(value interface{}) bool {
	if value == nil {
		return false
	}

	rv := reflect.ValueOf(value)

	switch value.(type) {
		case int, int8, int16, int32, int64:
			return rv.Int() != 0
		case uint, uint8, uint16, uint32, uint64:
			return rv.Uint() != 0
		case float32, float64:
			return rv.Float() != 0.0
		case string:
			return rv.String() != ""
		case bool:
			return rv.Bool()
		default:
			k := rv.Type().Kind()

			if k == reflect.Array || k == reflect.Slice || k == reflect.Map {
				return rv.Len() > 0
			}

			raise(errors.New("Unexpected type (" + k.String() + ") of value"), "ToBool")
	}
	return false
}

func ToBool(value interface{}) bool {
	return toBool(value)
}

func ToBoolSafe(value interface{}) (result bool, err error) {
	defer except(&err)
	result = toBool(value)
	return
}
