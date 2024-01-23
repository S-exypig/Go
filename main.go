package main

import (
	"demo/practise"
	"fmt"
)

func main() {
	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjsこん dk"
	s1ByteLen, s1RuneLen := practise.CountByteAndRune(s1)
	s2ByteLen, s2RuneLen := practise.CountByteAndRune(s2)
	fmt.Printf("string:%v, Byte=%v, Rune=%v\n", s1, s1ByteLen, s1RuneLen)
	fmt.Printf("string:%v, Byte=%v, Rune=%v\n", s2, s2ByteLen, s2RuneLen)
}
