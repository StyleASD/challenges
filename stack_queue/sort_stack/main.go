package sort_stack

import (
	"programming-problems/stack_queue"
)

/*
	Because we have loop through the data twice this algorithm will end up being an O(N^2) and O(N) space. This is the
	best we can do with the constraints we are given.
*/
func SortStack_Attempt1(stack stack_queue.Stack) stack_queue.Stack {
	tmpStack := stack_queue.NewStack()

	for !stack.IsEmpty() {
		pop := stack.Pop()

		// This loop allows us to compare the values and add any smaller values back to the first stack
		// This sorts the values by swapping values between stack until we have the largest at the bottom of tmpStack
		for !tmpStack.IsEmpty() && tmpStack.Peek().Data() > pop.Data() {
			stack.Push(tmpStack.Pop().Data())
		}
		// We can add the value we
		tmpStack.Push(pop.Data())
	}

	return *tmpStack
}
