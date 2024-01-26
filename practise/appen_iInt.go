package practise

import "fmt"

func AppendInt(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		fmt.Printf("newSlice:%v, len:%v, cap:%v\n", newSlice, len(newSlice), cap(newSlice))
		copy(newSlice, slice)
		slice = newSlice
		fmt.Printf("slice1:%v, len:%v, cap:%v\n", slice, len(slice), cap(slice))
	}
	slice = slice[0:n]
	fmt.Printf("slice2:%v, len:%v, cap:%v\n", slice, len(slice), cap(slice))
	copy(slice[m:n], data)
	return slice
}
