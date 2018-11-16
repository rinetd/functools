package functools

import (
	"testing"
	"reflect"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 0, 4, 5}, []int{1, 2, 4, 5}},
		{[]bool{true, true, true, true, true}, []bool{true, true, true, true, true}},
		{[]bool{true, false, true, false, true}, []bool{true, true, true}},
		{[]bool{false, false, false, false, false}, []bool{}},
		{[]float64{0.1, 5.2, 3.5, 1.005, 0.2}, []float64{0.1, 5.2, 3.5, 1.005, 0.2}},
		{[]float64{0.1, 0.0, 3.5, 1.005, 0.2}, []float64{0.1, 3.5, 1.005, 0.2}},
		{[]string{}, []string{}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "", "c"}, []string{"a", "c"}},
		{[]map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}},
			[]map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}},
		{[]map[int]int{{1: 1, 2: 2, 3: 3}, {}, {1: 1, 2: 2, 3: 3}},
			[]map[int]int{{1: 1, 2: 2, 3: 3}, {1: 1, 2: 2, 3: 3}}},
	}

	for _, c := range cases {
		got := Filter(c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Filter(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestFilterSafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
		e bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, false},
		{[]int{1, 2, 0, 4, 5}, []int{1, 2, 4, 5}, false},
		{map[int]int{1: 1, 2: 3, 3: 3}, nil, true},
		{true, nil, true},
	}

	for _, c := range cases {
		got, err := FilterSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("FilterSafe(%v) raised %v", c.in, err)
		}

		if err == nil && !reflect.DeepEqual(got, c.want) {
			t.Errorf("FilterSafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	correctFunc := func(n int) bool {
		return n % 2 == 0
	}

	cases := []struct {
		in []int
		want interface{}
	}{
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
		{[]int{3, 4, 5, 8, 9}, []int{4, 8}},
	}

	for _, c := range cases {
		got := FilterFunc(correctFunc, c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FilterFunc(correctFunc, %v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestFilterFuncSafe(t *testing.T) {
	correctFunc := func(n int) bool {
		return n % 2 == 0
	}

	incorrectFunc := func(a, b int) int {
		return a + b
	}

	cases := []struct {
		in []int
		want interface{}
	}{
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
		{[]int{3, 4, 5, 8, 9}, []int{4, 8}},
	}

	for _, c := range cases {
		got, err := FilterFuncSafe(correctFunc, c.in)

		if err != nil {
			t.Errorf("FilterFuncSafe(correctFunc, %v) raised %v", c.in, err)
		}

		if err == nil && !reflect.DeepEqual(got, c.want) {
			t.Errorf("FilterFuncSafe(correctFunc, %v) == %v want %v", c.in, got, c.want)
		}
	}

	_, err := FilterFuncSafe(incorrectFunc, cases[0].in)

	if err == nil {
		t.Errorf("FilterFuncSafe(incorrectFunc, %v) should raise type error here", cases[0].in)
	}
}
