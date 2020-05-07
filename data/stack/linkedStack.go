package stack

type LinkedStack struct {
	*stack
	root *linkedElement
}

type linkedElement struct {
	bottom *linkedElement
	value  interface{}
}

// Create the LinkedStack.
func NewLinkedStack() *LinkedStack {
	return new(LinkedStack).Init()
}

// Initialize the LinkedStack.
func (stack *LinkedStack) Init() *LinkedStack {
	stack.stack = initializeParent()
	stack.root = nil
	return stack
}

// Push a new item on the top.
func (stack *LinkedStack) Push(value interface{}) {
	newElement := linkedElement{
		bottom: stack.root,
		value:  value,
	}
	stack.root = &newElement
	stack.size++
}

// Top item returns. The objects returned will be removed.
func (stack *LinkedStack) Pop() interface{} {
	outObject := stack.root
	stack.root = outObject.bottom
	stack.size--
	return outObject.value
}

// Top item returns. Unlike 'Pop()', the top item returned will be maintained.
func (stack *LinkedStack) Peek() interface{} {
	outObject := stack.root
	return outObject.value
}
