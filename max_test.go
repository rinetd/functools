package functools

import (
	"testing"
	"reflect"
)

func TestMax(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
	}{
		{[]int{2, 4, 1, 5, 6, 9, 8, 0}, 9},
		{[]float64{1.1, 2.2, 5.5, 3.3}, 5.5},
		{[]string{"abc", "defg", "kl"}, "defg"},
		{[]rune{'a', 'b', 'c', 'd'}, 'd'},
		{[3]int{1, 2, 3}, 3},
		{[]map[int]int{{1: 1, 2: 2}, {3: 3}}, map[int]int{1: 1, 2: 2}},
	}

	for _, c := range cases {
		got := Max(c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Max(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestMaxSafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
		e bool
	}{
		{[]int{2, 4, 1, 5, 6, 9, 8, 0}, 9, false},
		{[]float64{1.1, 2.2, 5.5, 3.3}, 5.5, false},
		{[]string{"abc", "defg", "kl"}, "defg", false},
		{[]rune{'a', 'b', 'c', 'd'}, 'd', false},
		{[3]int{1, 2, 3}, 3, false},
		{[]map[int]int{{1: 1, 2: 2}, {3: 3}}, map[int]int{1: 1, 2: 2}, false},
		{[]int{}, 0, true},
		{1, 1, true},
	}

	for _, c := range cases {
		got, err := MaxSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("MaxSafe(%v) raised %v", c.in, err)
		}

		if err == nil && !reflect.DeepEqual(got, c.want) {
			t.Errorf("MaxSafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}
