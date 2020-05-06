package main

import (
	"GoStructure/data/stack"
	"log"
)

func main() {
	//exampleLinkedStack()
}

// Example of using LinkedStack
func exampleLinkedStack() {
	newInstance := stack.NewLinkedStack()
	newInstance.Push("Out")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	newInstance.Push("Last")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	newInstance.Push("Input")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	newInstance.Push("First")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	itemSize := newInstance.Size()
	for i := 0; i < itemSize; i++ {
		outValue := newInstance.Peek()
		log.Println("Peek value:", outValue, " Current size is : ", newInstance.Size())
	}

	for i := 0; i < itemSize; i++ {
		outValue := newInstance.Pop()
		log.Println("Pop value:", outValue, " Current size is : ", newInstance.Size())
		log.Println("IsEmpty:", newInstance.IsEmpty())
	}
}
