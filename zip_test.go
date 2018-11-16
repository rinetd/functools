package functools

import (
	"reflect"
	"testing"
)

type testZipStruct struct {
	a int
	b string
}

func TestZipMore(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want interface{}
	}{
		{[]interface{}{[][]int{{1, 4}, {2, 5}, {3, 6}}, []int{4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}}},
		{[]interface{}{[]int{1, 2, 3}, []int{4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}}},
	}

	for _, c := range cases {
		got := Zip(c.in...)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Zip(%v...) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestZip(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want interface{}
	}{
		{[]interface{}{[]int{1, 2, 3}, []int{4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}}},
		{[]interface{}{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[]interface{}{[]float64{1.1, 2.2}, []float64{3.3, 4.4}},
			[][]float64{{1.1, 3.3}, {2.2, 4.4}}},
		{[]interface{}{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8}},
			[][]int{{1, 6}, {2, 7}, {3, 8}, {4, 0}, {5, 0}}},
		{[]interface{}{[]int{}, []int{1, 2, 3}},
			[][]int{{0, 1}, {0, 2}, {0, 3}}},
		{[]interface{}{
			[]testZipStruct{
				{1, "a"},
				{2, "b"},
				{3, "c"},
			},
			[]testZipStruct{
				{4, "d"},
			},
		}, [][]testZipStruct{
			{testZipStruct{1, "a"}, testZipStruct{4, "d"}},
			{testZipStruct{2, "b"}, testZipStruct{0, ""}},
			{testZipStruct{3, "c"}, testZipStruct{0, ""}},
		}},
		{[]interface{}{[]string{"a1", "a2", "a3"}, []string{"b1", "b2", "b3"}},
			[][]string{{"a1", "b1"}, {"a2", "b2"}, {"a3", "b3"}}},
	}

	for _, c := range cases {
		got := Zip(c.in...)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Zip(%v...) == %v want %v", c.in, got, c.want)
		}
	}
}

func TestZipSafe(t *testing.T) {
	_, err := ZipSafe([]int{1, 2, 3})

	if err == nil {
		t.Error("ZipSafe should raise error '2 or more slices allowed'")
	}

	_, err = ZipSafe([3]int{1, 2, 3}, [3]int{4, 5, 6})

	if err == nil {
		t.Error("ZipSafe should raise error 'slices only allowed'")
	}

	_, err = ZipSafe([]int{1, 2, 3}, []int8{1, 2, 3})

	if err == nil {
		t.Error("ZipSafe should raise error 'params must be the same type'")
	}

	_, err = ZipSafe([]int{}, []int{})

	if err == nil {
		t.Error("ZipSafe should raise error 'at least one slice must have non-zero length'")
	}
}
