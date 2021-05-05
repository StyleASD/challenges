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

/*
	Attempt 1

	This is more complex than the above because we can no longer just add the carry. This solution works
	largely the same as the above with an additional step after adding the numbers to handle the carry

	The solution they use takes a different approach using objects to pass on the carry. It also once again uses recursion
	to to sum all the values. This time the recursion returns the partial sum with the object and this contains the carry
	the returned carry can then be added and passed to a new linked list
*/
func SumListForward_Attempt1(list1, list2 linked_lists.SinglyLinkedList) linked_lists.SinglyLinkedList {
	var lhs linked_lists.SinglyLinkedList
	var rhs linked_lists.SinglyLinkedList
	if list1.Length > list2.Length {
		lhs = list1
		rhs = list2
	} else {
		lhs = list2
		rhs = list1
	}

	buffer := lhs.Length - rhs.Length
	lhsHead := lhs.Head
	rhsHead := rhs.Head
	sum := []int{}

	for lhsHead != nil {
		if buffer > 0 {
			sum = append(sum, lhs.Head.Data)
		} else {
			addResult := lhsHead.Data + rhsHead.Data
			sum = append(sum, addResult)
		}
		lhsHead = lhsHead.Next
		rhsHead = rhsHead.Next
	}

	figures := []int{}
	carry := 0
	for i := len(sum) - 1; i != -1; i-- {
		val := sum[i] + carry

		if val > 10 && i == 0 {
			// If its the last value we need to handle it differently
			val++
			figures = append([]int{val % 10}, figures...)
			figures = append([]int{1}, figures...)
			break
		} else if val > 10 {
			carry = 1
		} else {
			carry = 0
		}

		// Modulo 10 will give us the value we want
		figures = append([]int{val % 10}, figures...)
	}

	return *linked_lists.NewSinglyLinkedList(figures...)
}
