package functools

import "testing"

func TestPartial(t *testing.T) {
	cases := []struct{
		f interface{}
		firstArgs []interface{}
		secondArgs []interface{}
		want interface{}
	}{
		{func(a, b int) int { return a + b }, []interface{}{2}, []interface{}{3}, 5},
		{func(a, b, c, d int) int { return a + b - c * d }, []interface{}{2, 3}, []interface{}{4, 5}, -15},
		{func(a, b, c string) string { return a + b + c }, []interface{}{"a", "b"}, []interface{}{"c"}, "abc"},
	}

	for _, c := range cases {
		got := Partial(c.f, c.firstArgs...)(c.secondArgs...)

		if got != c.want {
			t.Errorf("Partial(func, %v)(%v) == %v want %v", c.firstArgs, c.secondArgs, got, c.want)
		}
	}
}

func TestPartialSafe(t *testing.T) {
	// "a" is not a function
	_, err := PartialSafe("a", 1)

	if err == nil {
		t.Error("PartialSafe should raise error 'first argument isn't a function'")
	}

	// Mismatched types of func and parameter
	_, err = PartialSafe(func(a, b int) int { return a + b }, 1.2)

	if err == nil {
		t.Error("PartialSafe should raise error 'types of func and params aren't matched'")
	}

	// Passed more parameters than function allowes
	_, err = PartialSafe(func(a, b int) int { return a + b }, 1, 2, 3)

	if err == nil {
		t.Error("PartialSafe should raise error 'types of func and params aren't matched'")
	}
}
