package functools

import "testing"

var cmpCases = []struct {
	a    interface{}
	b    interface{}
	want int8
	e    bool
}{
	{nil, nil, 0, false},
	{nil, 0, -1, false},
	{0, nil, 1, false},
	{0, 0, 0, false},
	{1, 2, -1, false},
	{2, 1, 1, false},
	{3, 1, 1, false},
	{13.01, 13.02, -1, false},
	{0.0001, 0.0001, 0, false},
	{"abc", "cd", 1, false},
	{"abc", "cde", 0, false},
	{[]int{1, 2, 3}, []int{4, 5, 6}, 0, false},
	{[]float64{1.1, 2.2}, []float64{1.1, 2.2, 3.3}, -1, false},
	{false, false, 0, false},
	{true, false, 1, false},
	{false, true, -1, false},
	{1, 2.4, -1, true},
	{float64(1), 2.4, -1, false},
	{"a", 'a', 0, true},
	{map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{4: 4}, 1, false},
}

func TestCmpSafe(t *testing.T) {
	for _, c := range cmpCases {
		got, err := CmpSafe(c.a, c.b)

		if err != nil && !c.e {
			t.Errorf("CmpSafe(%v, %v) raised %v", c.a, c.b, err)
		}

		if err == nil && got != c.want {
			t.Errorf("CmpSafe(%v, %v) == %d want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestCmp(t *testing.T) {
	for _, c := range cmpCases {
		if !c.e {
			got := Cmp(c.a, c.b)

			if got != c.want {
				t.Errorf("Cmp(%v, %v) == %d want %d", c.a, c.b, got, c.want)
			}
		}
	}
}
