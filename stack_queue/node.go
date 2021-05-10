package stack_queue

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{data: data, next: nil}
}

func (n Node) Next() *Node {
	return n.next
}

func (n Node) Data() int {
	return n.data
}
