package functools

import (
	"reflect"
	"errors"
)

/*
'Partial' returns a new partial object which when called will behave like 'func'
called with the positional arguments args.

	Partial(func, args...) func(args...) interface{}
	PartialSafe(func, args...) (func(args...) interface{}, err)

 */

func partial(function interface{}, args ...interface{}) func(...interface{}) interface{} {
	fn := reflect.ValueOf(function)

	if fn.Kind() != reflect.Func {
		raise(errors.New("The first param should be a function"), "Partial")
	}

	reflectedArgs := make([]reflect.Value, 0, len(args))

	for _, arg := range args {
		reflectedArgs = append(reflectedArgs, reflect.ValueOf(arg))
	}

	if !verifyPartialFuncType(fn, reflectedArgs) {
		raise(errors.New("The types of function and params aren't matched"), "Partial")
	}

	partialedFunc := func(in ...interface{}) interface{} {
		params := make([]reflect.Value, 0, len(in) + len(reflectedArgs))
		params = append(params, reflectedArgs...)

		for _, v := range in {
			params = append(params, reflect.ValueOf(v))
		}

		return fn.Call(params[:])[0].Interface()
	}

	return partialedFunc
}

func verifyPartialFuncType(fn reflect.Value, args []reflect.Value) bool {
	if fn.Type().NumIn() <= len(args) {
		return false
	}

	for i := 0; i < len(args); i++ {
		if fn.Type().In(i) != args[i].Type() {
			return false
		}
	}

	return true
}

func Partial(function interface{}, args ...interface{}) func(...interface{}) interface{} {
	return partial(function, args...)
}

func PartialSafe(function interface{}, args ...interface{}) (result func(...interface{}) interface{}, err error) {
	defer except(&err)
	result = partial(function, args...)
	return
}
