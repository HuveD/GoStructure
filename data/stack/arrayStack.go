package stack

type ArrayStack struct {
	*Stack
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
	stack.Stack = initializeParent()
	stack.maxSize = capacity
	stack.itemArray = nil
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
