package functools

import (
	"reflect"
	"errors"
)

/*
'Zip' returns a slice of slices, where the i-th slice contains the i-th element
from each of the argument sequences. The returned slice has length of biggest slice
from passed arguments. If i-th element of passed slice is not exist - function puts
zero value of corresponding type instead of this element. Passed arguments must be
the same type.

	Zip(slices...) [][]interface{}
	ZipSafe(slices...) ([][]interface{}, err)
 */

func zip(slices ...interface{}) interface{} {
	if len(slices) < 2 {
		raise(errors.New("Function allows 2 or more slices only"), "Zip")
	}

	var sizes []int
	reflectedSlices := make([]reflect.Value, 0, len(slices))

	for _, slice := range slices {
		reflectedSlices = append(reflectedSlices, reflect.ValueOf(slice))
		sizes = append(sizes, reflect.ValueOf(slice).Len())
	}

	for _, rSlice := range reflectedSlices {
		if rSlice.Kind() != reflect.Slice {
			raise(errors.New("Params must be slices"), "Zip")
		}

		if rSlice.Type() != reflectedSlices[0].Type() {
			raise(errors.New("Params must be the same type"), "Zip")
		}
	}

	max_size := Max(sizes).(int)

	if max_size == 0 {
		raise(errors.New("At least one parameter must have non-zero length"), "Zip")
	}

	container := reflect.MakeSlice(reflectedSlices[0].Type(), len(reflectedSlices), len(reflectedSlices))
	out := reflect.MakeSlice(reflect.SliceOf(container.Type()), 0, max_size)

	for i := 0; i < max_size; i++ {
		for j, rSlice := range reflectedSlices {
			if i < rSlice.Len() {
				container.Index(j).Set(rSlice.Index(i))
			} else {
				container.Index(j).Set(reflect.Zero(rSlice.Type().Elem()))
			}
		}
		tmp := reflect.MakeSlice(container.Type(), container.Len(), container.Len())
		reflect.Copy(tmp, container)

		out = reflect.Append(out, tmp)
	}

	return out.Interface()
}

func Zip(slices ...interface{}) interface{} {
	return zip(slices...)
}

func ZipSafe(slices ...interface{}) (result interface{}, err error) {
	defer except(&err)
	result = zip(slices...)
	return
}
