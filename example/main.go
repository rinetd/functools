package main

import (
	"log"

	. "github.com/pytool/functools"
)

func ComposeT() {
	add := func(a, b int) int { return a + b }
	sub := func(a int) int { return a - 1 }
	mod := func(a int) bool { return a%2 == 0 }
	m := Compose(add, sub, mod)
	out := m(1, 2)
	log.Println(out)
}
func FilterT() {
	in := []int{1, 2, 0, 4, 5}
	out := FilterFunc(ToBool, in)
	log.Println(out)
}
func ApplyT() {
	// 返回Slice的长度
	len := func(s string) int { return len(s) }
	in := []string{"abc", "ab", "a"}
	out := Map(len, in)
	log.Println(out)
}
func PartialT() {
	sum := func(a, b int) int { return a * b }
	sum10 := Partial(sum, 10)
	result := sum10(10)
	log.Println(result)
}
func main() {
	ComposeT()
	PartialT()
	ApplyT()
	FilterT()
}
