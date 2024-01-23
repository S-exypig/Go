package user

import "fmt"

func PrintType() {
	a := 1
	b := 1.0
	c := "1"
	d := true
	e := []int{1, 2, 3}
	f := &a
	fmt.Printf("%v:%T\n", a, a)
	fmt.Printf("%v:%T\n", b, b)
	fmt.Printf("%v:%T\n", c, c)
	fmt.Printf("%v:%T\n", d, d)
	fmt.Printf("%v:%T\n", e, e)
	fmt.Printf("%v:%T\n", f, f)
}
