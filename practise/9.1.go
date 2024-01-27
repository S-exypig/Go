package practise

import (
	"container/list"
	"fmt"
)

func DoubleList() {
	listPtr := list.New()
	listPtr.PushBack(102)
	listPtr.PushFront(101)
	listPtr.PushBack(103)
	for e := listPtr.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
