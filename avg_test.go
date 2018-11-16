package functools

import "testing"

func TestAvg(t *testing.T) {
	cases := []struct {
		in interface{}
		want float64
	}{
		{[]int{1, 2, 3}, 2.0},
		{[]float32{2.3, 2.3, 2.3, 2.3}, 2.299999952316284},
		{[]float64{2.3, 2.3, 2.3, 2.3}, 2.3},
		{[]int8{1, 2, 3, 4, 5}, 3.0},
	}

	for _, c := range cases {
		got := Avg(c.in)

		if got != c.want {
			t.Errorf("Avg(%v) == %.4f want %.4f", c.in, got, c.want)
		}
	}
}

func TestAvgSafe(t *testing.T) {
	cases := []struct {
		in interface{}
		want float64
		e bool
	}{
		{[]int{1, 2, 3}, 2.0, false},
		{[]float32{2.3, 2.3, 2.3, 2.3}, 2.299999952316284, false},
		{[]float64{2.3, 2.3, 2.3, 2.3}, 2.3, false},
		{[]int8{1, 2, 3, 4, 5}, 3.0, false},
		{[]int{}, 0.0, true},
		{[]string{"a", "b", "c"}, 0.0, true},
	}

	for _, c := range cases {
		got, err := AvgSafe(c.in)

		if err != nil && !c.e {
			t.Errorf("AvgSafe(%v) raised %v", c.in, err)
		}

		if err == nil && got != c.want {
			t.Errorf("AvgSafe(%v) == %.4f want %.4f", c.in, got, c.want)
		}
	}
}
