package queue

type LinkedQueue struct {
	*queue
	head, tail *linkedElement
}

type linkedElement struct {
	pre, next *linkedElement
	value     interface{}
}

// Create the LinkedQueue.
func NewLinkedQueue() *LinkedQueue {
	return new(LinkedQueue).Init()
}

// Initialize the LinkedQueue.
func (queue *LinkedQueue) Init() *LinkedQueue {
	queue.queue = initializeParent()
	queue.head = nil
	queue.tail = nil
	return queue
}

// Add a new item on the top.
func (queue *LinkedQueue) Offer(value interface{}) {
	newElement := linkedElement{
		value: value,
	}
	originalTail := queue.tail
	queue.tail = &newElement
	if queue.size <= 0 {
		queue.head = &newElement
	} else {
		originalTail.pre = queue.tail
		queue.tail.next = originalTail
	}
	queue.size++
}

// Header item returns. The objects returned will be removed.
func (queue *LinkedQueue) Poll() interface{} {
	if queue.head == nil {
		return nil
	} else {
		//var outputValue interface{}
		outputValue := queue.head.value
		if queue.size > 1 {
			beforeItem := queue.head.pre
			queue.head = beforeItem
			queue.head.next = nil
		} else {
			queue.head = nil
			queue.tail = nil
		}
		queue.size--
		return outputValue
	}
}

// Top item returns. Unlike 'Poll()', the top item returned will be maintained.
func (queue *LinkedQueue) Peek() interface{} {
	outObject := queue.head
	if outObject == nil {
		return nil
	}
	return outObject.value
}
