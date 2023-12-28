package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Fuad")
	data.PushBack("Zaidan")
	data.PushBack("Najib")

	var head *list.Element = data.Front()
	head.Next().Next().Next()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
