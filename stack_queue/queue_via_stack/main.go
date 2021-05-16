package queue_via_stack

import (
	"programming-problems/stack_queue"
)

/*
	Attempt 1

	Queue is a first in first out structure which contrasts the the stack which is a last in first out structure. So to
	I need to maintain the order when popping using the second array. I can pop everything out of the first stack into the
	second stack until I get the to the last which is the value I want to return. I then add everything back to the first
	stack by performing this action in reverse which maintains the correct order.

	From the answer there is an optimisation we can do and that is to lazily move between stackForward and stackReverse
	as needed. Which I will implement below
*/
type MyQueue struct {
	stackForward stack_queue.Stack
	stackReverse stack_queue.Stack
}

func (queue *MyQueue) Pop() *stack_queue.Node {
	if queue.stackForward.IsEmpty() && queue.stackReverse.IsEmpty() {
		return nil
	}

	queue.reverseStacksForward()

	return queue.stackReverse.Pop()
}

func (queue *MyQueue) Peek() *stack_queue.Node {
	if queue.stackForward.IsEmpty() && queue.stackReverse.IsEmpty() {
		return nil
	}

	queue.reverseStacksForward()

	return queue.stackReverse.Peek()
}

// If the reverse is empty we can pop all the value onto the reverse stack
// While there are values in the stackReverse we can pop straight from there knowing they are in the correct order
// When it becomes empty again we fetch more of the values form stackForward
func (queue *MyQueue) reverseStacksForward()  {
	if queue.stackReverse.IsEmpty() {
		for !queue.stackForward.IsEmpty() {
			queue.stackReverse.Push(queue.stackForward.Pop().Data()) // Add to Stack to in reverse order
		}
	}
}

func (queue *MyQueue) Push(data int) {
	queue.stackForward.Push(data) // We can push to stackForward knowing that order is maintained by stackReverse
}
