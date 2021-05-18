package stack_queue

// First in first out
type Stack struct {
	Head *Node
	last *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s Stack) Pop() *Node {
	if s.Head == nil {
		return nil
	}

	current := s.Head
	s.Head = current.Next()
	current.next = nil

	return current
}

func (s Stack) Push(data int) {
	node := NewNode(data)
	if s.Head == nil {
		s.Head = node
		s.last = node
		return // Exit function
	}

	s.last.next = node
	s.last = node
}

func (s Stack) Peek() *Node {
	if s.Head == nil {
		return nil
	}

	return s.Head
}

func (s Stack) IsEmpty() bool {
	return s.Head == nil
}
