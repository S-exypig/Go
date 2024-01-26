package practise

import "fmt"

func Varags(nums ...int) {
	if len(nums) == 0 {
		return
	}
	for i, v := range nums {
		fmt.Printf("i:%v v: %v\n", i, v)
	}
}
