package main

func main() {

}

/*
	Attempt 1

	This works in O(1) time by keeping track of the min in each node. This is not as space efficient as it could be I
	think this could be made more efficient than this implementation because the memory address on all nodes track the
	address to the the min value at the time it was pushed.
 */

type StackMin struct {
	HEAD *Node
	TAIL *Node
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
	Min  *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

func (stack *StackMin) Push(data int) (ok bool) {
	node := NewNode(data)

	if stack.TAIL != nil {
		if stack.TAIL.Min.Data <= data {
			node.Min = node
		} else {
			node.Min = stack.TAIL.Min
		}

		node.Prev = stack.TAIL

		stack.TAIL.Next = node
		stack.TAIL = node
	} else {
		stack.HEAD = node
		stack.TAIL = node
	}

	return true
}

func (stack *StackMin) Pop() Node {
	node := stack.TAIL
	stack.TAIL = node.Prev
	stack.TAIL.Next = nil
	node.Prev = nil

	return *node
}

func (stack *StackMin) Min() Node {
	return *stack.TAIL.Min
}
