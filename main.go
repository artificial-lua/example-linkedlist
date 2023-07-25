package main

import (
	"fmt"

	"github.com/artificial-lua/example-linkedlist/linkedlist"
)

func main() {
	sampleSingleLinkedList := linkedlist.NewSingleLinkedList()
	sampleSingleLinkedList.Append(1)
	sampleSingleLinkedList.Append(2)
	sampleSingleLinkedList.Append(3)
	sampleSingleLinkedList.Append(4)
	sampleSingleLinkedList.Append(5)
	node, err := sampleSingleLinkedList.SearchByIndex(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(node.GetValue())
	fmt.Println(sampleSingleLinkedList)

	sampleDoubleLinkedList := linkedlist.NewDoubleLinkedList()
	sampleDoubleLinkedList.Append("1")
	sampleDoubleLinkedList.Append("a")
	sampleDoubleLinkedList.Append(" ")
	sampleDoubleLinkedList.Append("J")
	fmt.Println(sampleDoubleLinkedList)
	sampleDoubleLinkedList.Pop(2)
	fmt.Println(sampleDoubleLinkedList)
	sampleDoubleLinkedList.Update(1, "b")
	fmt.Println(sampleDoubleLinkedList)
	sampleDoubleLinkedList.Prepend("0")
	sampleDoubleLinkedList.Append("K")
	sampleDoubleLinkedList.Insert(3, "c")
	sampleDoubleLinkedList.Insert(2, sampleSingleLinkedList)
	fmt.Println(sampleDoubleLinkedList)
}
