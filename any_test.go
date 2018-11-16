package functools

import "testing"

func TestAny(t *testing.T) {
	cases := []struct {
		in interface{}
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 0, 4, 5}, true},
		{[]int{0, 0, 0, 0, 1, 0}, true},
		{[]int{0, 0, 0, 0, 0, 0}, false},
		{[]bool{true, true, true, true, true}, true},
		{[]bool{true, true, true, false, true}, true},
		{[]bool{false, false, false, false, false}, false},
		{[]float64{0.1, 5.2, 3.5, 1.005, 0.2}, true},
		{[]float64{0.1, 0.0, 3.5, 1.005, 0.2}, true},
		{[]float64{0.0, 0.0, 0.0}, false},
		{[]string{}, false},
		{[]string{"a", "b", "c"}, true},
		{[]string{"a", "", "c"}, true},
		{[]string{"", "", ""}, false},
		{[]map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}, true},
		{[]map[int]int{{1: 1, 2: 2, 3: 3}, {}}, true},
		{[]map[int]int{{}, {}, {}}, false},
	}

	for _, c := range cases {
		got := Any(c.in)

		if got != c.want {
			t.Errorf("Any(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestAnySafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want bool
		e bool
	}{
		{[]int{1, 2, 3, 4, 5}, true, false},
		{[]int{1, 2, 0, 4, 5}, true, false},
		{[]int{0, 0, 0}, false, false},
		{map[int]int{1: 1, 2: 3, 3: 3}, false, true},
		{true, false, true},
	}

	for _, c := range cases {
		got, err := AnySafe(c.in)

		if err != nil && !c.e {
			t.Errorf("AnySafe(%v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("AnySafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestAnyFunc(t *testing.T) {
	correctFunc := func(n int) bool {
		return n % 2 == 0
	}

	cases := []struct {
		in []int
		want bool
	}{
		{[]int{1, 4, 3, 5, 7}, true},
		{[]int{1, 3, 5, 7, 9}, false},
	}

	for _, c := range cases {
		got := AnyFunc(correctFunc, c.in)

		if got != c.want {
			t.Errorf("AnyFunc(correctFunc, %v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestAnyFuncSafe(t *testing.T) {
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
		{[]int{1, 4, 3, 5, 7}, true},
		{[]int{1, 3, 5, 7, 9}, false},
	}

	for _, c := range cases {
		got, err := AnyFuncSafe(correctFunc, c.in)

		if err != nil {
			t.Errorf("AnyFuncSafe(correctFunc, %v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("AnyFuncSafe(correctFunc, %v) == %v want %v", c.in, got, c.want)
		}
	}

	_, err := AnyFuncSafe(incorrectFunc, cases[0].in)

	if err == nil {
		t.Errorf("AnyFuncSafe(incorrectFunc, %v) should raise type error here", cases[0].in)
	}
}
