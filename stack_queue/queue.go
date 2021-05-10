package stack_queue

// Last in first out
type QueueImpl struct {
	Head *Node
	Last *Node
}

func (q QueueImpl) add(data int) {
	node := NewNode(data)

	if q.Head == nil {
		q.Head = node
		q.Last = node
		return // Exit function
	}

	node.next = q.Head
	q.Head = node

}

func (q QueueImpl) remove() *Node {
	if q.Head == nil {
		return nil
	}

	current := q.Head
	q.Head = current.Next()
	current.next = nil

	return current

}

func (q QueueImpl) peek() Node {
	return *q.Head
}

func (q QueueImpl) isEmpty() bool {
	return q.Head == nil
}
