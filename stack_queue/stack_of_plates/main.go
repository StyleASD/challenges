package stack_of_plates

import (
	"fmt"
)

/*
	Attempt 1

	This implementation seems similar to that in the example answer the main differences come form nuances in the go language.

	There is also the consideration of whether I should keep the capacity in the PopAt function consistent by shifting all
	the values left. I choose to no keep the capacities the same in order to keep the PopAt function working in O(1) time.
*/
type SetOfStacks struct {
	stacks   []Stack
	capacity int
	stackNo  int
}

func NewSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{stacks: make([]Stack, 1), capacity: capacity, stackNo: 0}
}

type Stack struct {
	Tail   *Node
	Length int
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func (s *SetOfStacks) Pop() *Node {
	stack := s.stacks[s.stackNo]

	if stack.Length < 1 {
		return nil
	}

	pop := stack.Tail
	stack.Tail = pop.Prev
	pop.Prev = nil
	stack.Length--

	if stack.Length == 0 && len(s.stacks) > 1 {
		s.stacks = s.stacks[:len(s.stacks)-1] // Remove stack if we no longer need it 
		s.stackNo--
	}

	return pop
}

func (s *SetOfStacks) Push(data int) {
	// Check to see if the current stack is at capacity if it is create a new one
	if s.stacks[s.stackNo].Length == s.capacity {
		s.stacks = append(s.stacks, Stack{})
		s.stackNo++
	}

	stack := s.stacks[s.stackNo]

	// Create the Node and add it to the linked list
	node := &Node{Data: data, Prev: stack.Tail}
	stack.Tail.Next = node
	stack.Tail = node
	stack.Length++
}

func (s *SetOfStacks) PopAt(index int) (*Node, error) {
	if s.stackNo > index && index < 0 {
		return nil, fmt.Errorf("Index %s out of range %s", index, s.stacks)
	}

	stack := s.stacks[index]
	pop := stack.Tail
	stack.Tail = pop.Prev
	pop.Prev = nil
	stack.Length--

	return pop, nil
}
