package stack

import "errors"

type ArrayStack struct {
	*stack
	itemArray []arrayElement
	maxSize   int
	pointer   int
}

type arrayElement struct {
	index int
	value interface{}
}

// Create the ArrayStack.
func NewArrayStack(capacity int) *ArrayStack {
	return new(ArrayStack).Init(capacity)
}

// Initialize the ArrayStack.
func (stack *ArrayStack) Init(capacity int) *ArrayStack {
	stack.stack = initializeParent()
	stack.itemArray = make([]arrayElement, capacity)
	stack.maxSize = capacity
	stack.pointer = 0
	return stack
}

// Returns the maximum item size.
func (stack *ArrayStack) Capacity() int {
	return stack.maxSize
}

// Check if the Stack is full.
func (stack *ArrayStack) IsFull() bool {
	return stack.Capacity() == stack.size
}

// Push a new item on the top.
func (stack *ArrayStack) Push(value interface{}) error {
	newElement := arrayElement{
		index: stack.size,
		value: value,
	}
	endIndex := newElement.index
	if endIndex >= stack.Capacity() {
		return errors.New("capacity exceeded")
	} else {
		stack.itemArray[endIndex] = newElement
		stack.pointer = endIndex
		stack.size++
		return nil
	}
}

// Top item returns. The objects returned will be removed.
func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	} else {
		outObject := stack.itemArray[stack.pointer]
		stack.itemArray[stack.pointer] = arrayElement{}
		stack.pointer--
		stack.size--
		return outObject.value
	}
}

// Top item returns. Unlike 'Pop()', the top item returned will be maintained.
func (stack *ArrayStack) Peek() interface{} {
	outObject := stack.itemArray[stack.pointer]
	return outObject.value
}
