package functools

import "testing"

func TestReduce(t *testing.T) {
	cases := []struct{
		f interface{}
		in interface{}
		acc interface{}
		want interface{}
	}{
		{func(a, b int) int { return a + b }, []int{1, 2, 3, 4, 5}, 0, 15},
		{func(a, b string) string { return a + b }, []string{"a", "b", "c"}, "", "abc"},
		{func(a, b float64) float64 { return a * b }, []float64{0.1, 2.5, 9.3}, 1.0, 2.325},
		{func(a, b int) int { return a + b }, []int{}, 0, 0},
		{func(a, b interface{}) bool {return ToBool(a) && ToBool(b) }, []bool{true, false, true}, true, false},
	}

	for _, c := range cases {
		got := Reduce(c.f, c.in, c.acc)

		if got != c.want {
			t.Errorf("Reduce(func, %v, %v) == %v want %v", c.in, c.acc, got, c.want)
		}
	}
}

func TestReduceSafe(t *testing.T) {
	// Different types of func arguments
	_, err := ReduceSafe(func(a float64, b int) float64 { return a + float64(b) }, []int{1, 2, 3}, 0.0)

	if err == nil {
		t.Error("ReduceSafe should raise error of type of func here")
	}

	// Wrong number of func arguments
	_, err = ReduceSafe(func(a, b, c int) int { return a + b + c }, []int{1, 2, 3}, 0)

	if err == nil {
		t.Error("ReduceSafe should raise error of type of func here")
	}

	// Type of accumulator is not the same as type of func
	_, err = ReduceSafe(func(a, b float64) float64 {return a + b }, []float64{1, 2}, 0)

	if err == nil {
		t.Error("ReduceSafe should raise error of type of accumulator here")
	}

	// Accumulator is nil
	_, err = ReduceSafe(func(a, b int) int { return a + b }, []int{1, 2}, nil)

	if err == nil {
		t.Error("ReduceSafe should raise error of type of accumulator here")
	}
}
