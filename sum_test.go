package functools

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
	}{
		{[]int{2, 4, 1, 5, 6, 9, 8, 0}, 35},
		{[]float64{1.1, 2.2, 5.5, 3.3}, 12.100000000000001},
		{[]int16{1, 2, 3}, int16(6)},
		{[]float32{1.0, 1.0, 1.0}, float32(3.0)},
	}

	for _, c := range cases {
		got := Sum(c.in)

		if got != c.want {
			t.Errorf("Sum(%v) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestSumSafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want interface{}
		e bool
	}{
		{[]int{2, 4, 1, 5, 6, 9, 8, 0}, 35, false},
		{[]float64{1.1, 2.2, 5.5, 3.3}, 12.100000000000001, false},
		{[]int16{1, 2, 3}, int16(6), false},
		{[]float32{1.0, 1.0, 1.0}, float32(3.0), false},
		{[]string{"a", "b"}, 0.0, true},
		{[]int{}, 0.0, true},
	}

	for _, c := range cases {
		got, err := SumSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("SumSafe(%v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("SumSafe(%v) == %v want %v", c.in, got, c.want)
		}
	}
}
