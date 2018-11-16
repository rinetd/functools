package functools

import "testing"

type unexpectedType struct {
	field int
}

var boolCases = []struct {
	in interface{}
	want interface{}
	e bool
}{
	{0, false, false},
	{1, true, false},
	{-1, true, false},
	{0.0, false, false},
	{0.1, true, false},
	{-0.1, true, false},
	{uint(2), true, false},
	{uint8(0), false, false},
	{uint16(1), true, false},
	{uint32(0), false, false},
	{uint64(5), true, false},
	{int8(0), false, false},
	{int16(1), true, false},
	{int32(0), false, false},
	{int64(5), true, false},
	{"abc", true, false},
	{"", false, false},
	{[]int{}, false, false},
	{[]int(nil), false, false},
	{[]int{1, 2, 3}, true, false},
	{[3]int{}, true, false},
	{[]float64{}, false, false},
	{[]float64{1.0, 2.1, 3.5}, true, false},
	{[]string{}, false, false},
	{[]string{"a", "b", "c"}, true, false},
	{[]rune{}, false, false},
	{[]rune{'a', 'b', 'c'}, true, false},
	{float32(0.001), true, false},
	{nil, false, false},
	{'a', true, false},
	{byte('a'), true, false},
	{make(map[string]int), false, false},
	{map[string]int{"a": 1}, true, false},
	{true, true, false},
	{false, false, false},
	{unexpectedType{1}, false, true},
}

func TestToBoolSafe(t *testing.T) {
	for _, c := range boolCases {
		got, err := ToBoolSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("ToBoolSafe(%v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("ToBoolSafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestToBool(t *testing.T) {
	for _, c := range boolCases {
		if !c.e {
			got := ToBool(c.in)

			if got != c.want {
				t.Errorf("ToBool(%v) == %v want %v", c.in, got, c.want)
			}
		}
	}
}
