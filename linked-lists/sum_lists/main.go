package main

import (
	linked_lists "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	Calculating the reverse order is easier than the front because we can work out the remainder as we move along.
	The algorithm I came up with is an O(N) in time I also think there it could be tidied up.

	Moving the algorithm when the like list is not in reverse I need to figure out how I will handle the remainder.

	From the answer we can also do this with recursion I think this would use the same amount of space as my algorithm
	which I think is O(N).

*/
func SumListReverse_Attempt1(list1, list2 linked_lists.SinglyLinkedList) linked_lists.SinglyLinkedList {
	lhsHead := list1.Head
	rhsHead := list2.Head

	sum := []int{}
	carry := 0

	for lhsHead != nil {
			// Add lhs to rhs
			addResult := lhsHead.Data + rhsHead.Data + carry

			// We can work out the carry by testing if the number is greater than 10
			if addResult > 10 && lhsHead.Next == nil {
				// If its the last value we need to handle it differently
				addResult++
				sum = append(sum, addResult%10)
				sum = append(sum, 1)
				break
			} else if addResult > 10 {
				carry = 1
			} else {
				carry = 0
			}

			// Modulo 10 will give us the value we want
			sum = append(sum, addResult%10)
			lhsHead = lhsHead.Next
			rhsHead = rhsHead.Next
	}

	return *linked_lists.NewSinglyLinkedList(sum...)
}
