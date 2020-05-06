package stack

// Queue struct
type Stack struct {
	size int
}

func initializeParent() *Stack {
	parent := new(Stack)
	parent.size = 0
	return parent
}

// Check if the Stack is empty.
func (stack *Stack) IsEmpty() bool {
	return stack.size <= 0
}

// Get Item Size.
func (stack *Stack) Size() int {
	return stack.size
}
