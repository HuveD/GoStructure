package queue

// Queue struct
type queue struct {
	size int
}

func initializeParent() *queue {
	parent := new(queue)
	parent.size = 0
	return parent
}

// Check if the Stack is empty.
func (stack *queue) IsEmpty() bool {
	return stack.size <= 0
}

// Get Item Size.
func (stack *queue) Size() int {
	return stack.size
}
