package functools

import (
	"strconv"
	"strings"
	"testing"
)

func TestCompose(t *testing.T) {
	cases := []struct {
		fs   []interface{}
		in   []interface{}
		want interface{}
	}{
		{fs: []interface{}{
			func(a, b int) int { return a + b },
			func(a int) int { return a - 1 },
			func(a int) bool { return a%2 == 0 }},
			in:   []interface{}{1, 2},
			want: true,
		},
		{[]interface{}{
			func(a, b, c string) string { return a + b + c },
			func(a string) string { return strings.ToUpper(a) },
			func(a string) int {
				v, _ := strconv.Atoi(a)
				return v
			},
			ToBool,
		}, []interface{}{"1", "2", "3"}, true},
		{[]interface{}{
			All,
			func(a bool) string {
				if a {
					return "OK"
				}
				return "Not OK"
			},
		}, []interface{}{[]int{1, 2, 3, 4, 0, 5, 6}}, "Not OK"},
	}

	for _, c := range cases {
		run := Compose(c.fs...)
		run(c.in...)
		got := Compose(c.fs...)(c.in...)

		if got != c.want {
			t.Errorf("Compose(fs...)(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestComposeSafe(t *testing.T) {
	_, err := ComposeSafe(All, Any, 4, ToBool)

	if err == nil {
		t.Error("ComposeSafe should raise error 'param isn't a function'")
	}

	_, err = ComposeSafe(All, func(a int) int { return a + 1 }, ToBool)

	if err == nil {
		t.Error("ComposeSafe should raise error 'functions cannot be piped'")
	}

	_, err = ComposeSafe(func(a, b int) int { return a + b }, func(a, b int) int { return a + b })

	if err == nil {
		t.Error("ComposeSafe should raise error 'functions cannot be piped'")
	}
}
