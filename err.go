package functools

import "fmt"

// This way of error handling based on
// https://github.com/reusee/socks5hs/blob/master/err.go
// https://github.com/choleraehyq/gofunctools/blob/master/functools/err.go

type customErr struct {
	pkg string
	info string
	err error
}

func (e *customErr) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s : %s\n%v", e.pkg, e.info, e.err)
	}
	return fmt.Sprintf("%s: %s\n", e.pkg, e.info)
}

func generateErr(e error, format string, args ...interface{}) *customErr {
	if len(args) > 0 {
		return &customErr{
			pkg: "go-built-in",
			info: fmt.Sprintf(format, args...),
			err: e,
		}
	}
	return &customErr{
		pkg: "go-built-in",
		info: format,
		err: e,
	}
}

func raise(e error, format string, args ...interface{}) {
	if e != nil {
		panic(generateErr(e, format, args...))
	}
}

func except(e *error) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*e = err
		} else {
			panic(r)
		}
	}
}
