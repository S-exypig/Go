package practise

import "fmt"

func ForCharacter() {
	/* for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	*/
	/* 	origStr := "GGGGGGGGGGGGGGGGGGGGGGGGG"
	   	for i := 1; i <= 25; i++ {
	   		newStr := origStr[0:i]
	   		fmt.Printf("%v\n", newStr)
	   	} */
	origStr := "G"
	for i := 1; i <= 25; i++ {
		fmt.Printf("%v\n", origStr)
		origStr += "G"
	}
}
