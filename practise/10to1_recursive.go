package practise

import "fmt"

func Recursive(num int) {
	if num == 11 {
		return
	}
	fmt.Printf("%v\n", num)
	Recursive(num + 1)
}
