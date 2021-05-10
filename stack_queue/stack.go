package stack_queue

// First in first out
type Stack struct {
	Head *Node
	last *Node
}

func (s Stack) pop() *Node {
	if s.Head == nil {
		return nil
	}

	current := s.Head
	s.Head = current.Next()
	current.next = nil

	return current
}

func (s Stack) push(data int) {
	node := NewNode(data)
	if s.Head == nil {
		s.Head = node
		s.last = node
		return // Exit function
	}

	s.last.next = node
	s.last = node
}

func (s Stack) peek() *Node {
	if s.Head == nil {
		return nil
	}

	return s.Head
}

func (s Stack) isEmpty() bool {
	return s.Head == nil
}
