package practise

import "fmt"

func MagnifySlice(s []int, factor int) []int {
	defer fmt.Printf("len(s):%v, cap(s):%v\n", len(s), cap(s))
	res := make([]int, len(s)*factor)
	defer fmt.Printf("len(res):%v, cap(res):%v\n", len(res), cap(res))
	copy(res, s)
	return res
}
