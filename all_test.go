package functools

import "testing"

func TestAll(t *testing.T) {
	cases := []struct {
		in interface{}
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 0, 4, 5}, false},
		{[]bool{true, true, true, true, true}, true},
		{[]bool{true, true, true, false, true}, false},
		{[]bool{false, false, false, false, false}, false},
		{[]float64{0.1, 5.2, 3.5, 1.005, 0.2}, true},
		{[]float64{0.1, 0.0, 3.5, 1.005, 0.2}, false},
		{[]string{}, true},
		{[]string{"a", "b", "c"}, true},
		{[]string{"a", "", "c"}, false},
		{[]map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}, true},
		{[]map[int]int{{1: 1, 2: 2, 3: 3}, {}, {1: 1, 2: 2, 3: 3}}, false},
	}

	for _, c := range cases {
		got := All(c.in)

		if got != c.want {
			t.Errorf("All(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestAllSafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want bool
		e bool
	}{
		{[]int{1, 2, 3, 4, 5}, true, false},
		{[]int{1, 2, 0, 4, 5}, false, false},
		{map[int]int{1: 1, 2: 3, 3: 3}, false, true},
		{true, false, true},
	}

	for _, c := range cases {
		got, err := AllSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("AllSafe(%v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("AllSafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestAllFunc(t *testing.T) {
	correctFunc := func(n int) bool {
		return n % 2 == 0
	}

	cases := []struct {
		in []int
		want bool
	}{
		{[]int{2, 4, 6, 8, 10}, true},
		{[]int{2, 4, 5, 8, 10}, false},
	}

	for _, c := range cases {
		got := AllFunc(correctFunc, c.in)

		if got != c.want {
			t.Errorf("AllFunc(correctFunc, %v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestAllFuncSafe(t *testing.T) {
	correctFunc := func(n int) bool {
		return n % 2 == 0
	}

	incorrectFunc := func(a, b int) int {
		return a + b
	}

	cases := []struct {
		in []int
		want bool
	}{
		{[]int{2, 4, 6, 8, 10}, true},
		{[]int{2, 4, 5, 8, 10}, false},
	}

	for _, c := range cases {
		got, err := AllFuncSafe(correctFunc, c.in)

		if err != nil {
			t.Errorf("AllFuncSafe(correctFunc, %v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("AllFuncSafe(correctFunc, %v) == %v want %v", c.in, got, c.want)
		}
	}

	_, err := AllFuncSafe(incorrectFunc, cases[0].in)

	if err == nil {
		t.Errorf("AllFuncSafe(incorrectFunc, %v) should raise type error here", cases[0].in)
	}
}
