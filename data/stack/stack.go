package stack

// Stack struct
type stack struct {
	size int
}

func initializeParent() *stack {
	parent := new(stack)
	parent.size = 0
	return parent
}

// Check if the Stack is empty.
func (stack *stack) IsEmpty() bool {
	return stack.size <= 0
}

// Get Item Size.
func (stack *stack) Size() int {
	return stack.size
}
