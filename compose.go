package functools

import (
	"reflect"
	"errors"
	"strconv"
)

/*
'Compose' allows to pipe unlimited number of functions to one composed 'func'.

	Compose(fs...) func(args...) interface{}
	ComposeSafe(fs...) (func(args...) interface{}, err)
 */

func compose(functions ...interface{}) func(...interface{}) interface{} {
	for i, fn := range functions {
		if reflect.ValueOf(fn).Kind() != reflect.Func {
			raise(errors.New("Param " + strconv.Itoa(i) + " isn't a function"), "Compose")
		}
	}

	for i := 0; i < len(functions) - 1; i++ {
		fnA := reflect.ValueOf(functions[i])
		fnB := reflect.ValueOf(functions[i + 1])

		if !canPipe(fnA, fnB) {
			raise(errors.New("Function " + strconv.Itoa(i) + " and " +
				strconv.Itoa(i + 1) + " cannot be piped"), "Compose")
		}
	}

	composedFunc := func(in ...interface{}) interface{} {
		params := make([]reflect.Value, 0, len(in))

		for _, v := range in {
			params = append(params, reflect.ValueOf(v))
		}

		for i := 0; i < len(functions); i++ {
			params = reflect.ValueOf(functions[i]).Call(params[:])
		}

		return params[0].Interface()
	}

	return composedFunc
}

func canPipe(fnA, fnB reflect.Value) bool {
	if fnA.Type().NumOut() != fnB.Type().NumIn() {
		return false
	}

	for i := 0; i < fnA.Type().NumOut(); i++ {
		if fnA.Type().Out(i) != fnB.Type().In(i) && fnB.Type().In(i).Kind() != reflect.Interface {
			return false
		}
	}

	return true
}

func Compose(functions ...interface{}) func(...interface{}) interface{} {
	return compose(functions...)
}

func ComposeSafe(functions ...interface{}) (result func(...interface{}) interface{}, err error) {
	defer except(&err)
	result = compose(functions...)
	return
}
