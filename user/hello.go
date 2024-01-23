package user

import "fmt"

func Hello() string {
	a, b := "Hello", "World"
	fmt.Printf("%s\n", a+b)
	return "Hello"
}
