package functools

import (
	"testing"
	"reflect"
)

func TestMap(t *testing.T) {
	cases := []struct{
		f interface{}
		in interface{}
		want interface{}
	}{
		{func(s string) int { return len(s) }, []string{"abc", "ab", "a"}, []int{3, 2, 1}},
		{ToBool, []int{1, 2, 3, 0, 5, 6, 0, 8}, []bool{true, true, true, false, true, true, false, true}},
		{func(n int) int { return n % 2 }, []int{1, 2, 3, 4, 5, 6}, []int{1, 0, 1, 0, 1, 0}},
		{ToBool, []int{}, []bool{}},
	}

	for _, c := range cases {
		got := Map(c.f, c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Map(func, %v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestMapSafe(t *testing.T) {
	_, err := MapSafe(ToBool, 123)

	if err == nil {
		t.Errorf("FilterFuncSafe(ToBool, %v) should raise type error here", 123)
	}

	_, err = MapSafe(func (a, b int) bool { return a == b }, []int{1, 2, 3})

	if err == nil {
		t.Errorf("FilterFuncSafe(func, %v) should raise incorrect func error here", []int{1, 2, 3})
	}
}
