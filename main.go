package main

import (
	"GoStructure/data/queue"
	"GoStructure/data/stack"
	"log"
)

func main() {
	//exampleLinkedStack()
	//exampleArrayStack()
	//exampleLinkedQueue()
}

//region Stack example

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

// Example of using ArrayStack
func exampleArrayStack() {
	newInstance := stack.NewArrayStack(4)

	err := newInstance.Push("Out")
	if err == nil {
		log.Printf("Added new one, Current size is %d\n", newInstance.Size())
	} else {
		log.Println(err)
	}

	err = newInstance.Push("Last")
	if err == nil {
		log.Printf("Added new one, Current size is %d\n", newInstance.Size())
	} else {
		log.Println(err)
	}

	err = newInstance.Push("Input")
	if err == nil {
		log.Printf("Added new one, Current size is %d\n", newInstance.Size())
	} else {
		log.Println(err)
	}

	err = newInstance.Push("First")
	if err == nil {
		log.Printf("Added new one, Current size is %d\n", newInstance.Size())
	} else {
		log.Println(err)
	}

	err = newInstance.Push("Firs232323t")
	if err == nil {
		log.Printf("Added new one, Current size is %d\n", newInstance.Size())
	} else {
		log.Println(err)
	}

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

//endregion Stack example

//region Queue example

// Example of using LinkedQueue
func exampleLinkedQueue() {
	newInstance := queue.NewLinkedQueue()

	newInstance.Offer("First")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	newInstance.Offer("Input")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	newInstance.Offer("First")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	newInstance.Offer("Out")
	log.Printf("Added new one, Current size is %d\n", newInstance.Size())

	itemSize := newInstance.Size()
	for i := 0; i < itemSize; i++ {
		outValue := newInstance.Peek()
		log.Println("Peek value:", outValue, " Current size is : ", newInstance.Size())
	}

	for i := 0; i < itemSize; i++ {
		outValue := newInstance.Poll()
		log.Println("Poll value:", outValue, " Current size is : ", newInstance.Size())
		log.Println("IsEmpty:", newInstance.IsEmpty())
	}
}

//endregion Queue example
