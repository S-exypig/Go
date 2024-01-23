package practise

import "fmt"

func ForLoop() {
	for i := 0; i < 15; i++ {
		fmt.Printf("i: %v\n", i)
	}

	var i int = 0
START:
	fmt.Printf("i: %v\n", i)
	i++
	if i < 15 {
		goto START
	}
}
