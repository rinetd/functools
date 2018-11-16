package functools

import (
	"testing"
	"reflect"
)

func TestMin(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
	}{
		{[]int{2, 4, 1, 5, 6, 9, 8, 0}, 0},
		{[]float64{1.1, 2.2, 5.5, 3.3}, 1.1},
		{[]string{"abc", "defg", "kl"}, "kl"},
		{[]rune{'a', 'b', 'c', 'd'}, 'a'},
		{[3]int{1, 2, 3}, 1},
		{[]map[int]int{{1: 1, 2: 2}, {3: 3}}, map[int]int{3: 3}},
	}

	for _, c := range cases {
		got := Min(c.in)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Min(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestMinSafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
		e bool
	}{
		{[]int{2, 4, 1, 5, 6, 9, 8, 0}, 0, false},
		{[]float64{1.1, 2.2, 5.5, 3.3}, 1.1, false},
		{[]string{"abc", "defg", "kl"}, "kl", false},
		{[]rune{'a', 'b', 'c', 'd'}, 'a', false},
		{[3]int{1, 2, 3}, 1, false},
		{[]map[int]int{{1: 1, 2: 2}, {3: 3}}, map[int]int{3: 3}, false},
		{[]int{}, 0, true},
		{1, 1, true},
	}

	for _, c := range cases {
		got, err := MinSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("MinSafe(%v) raised %v", c.in, err)
		}

		if err == nil && !reflect.DeepEqual(got, c.want) {
			t.Errorf("MinSafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}
