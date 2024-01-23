package practise

import (
	"fmt"
	"strconv"
)

func ErrorChange() {
	var origStr string = "123"
	// var an int
	var newSrt string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(origStr)
	if err != nil {
		fmt.Printf("origStr %s is not an integer - exiting with error\n", origStr)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newSrt = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newSrt)
}
