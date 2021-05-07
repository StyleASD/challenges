package main

import (
	linked_lists "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	I came up with this fairly simply. A palindrome basically requires that the first half is equal to that of the second
	half. So I thought if we reverse the first half by placing the data in an array we can and then iterate over the array
	to check the latter half values. This ends up working in 0(N) time but has an O(N/2) space requirement for the additional
	array.

	The solution has two options the examples also dont have a length solution

	1) First solution is to reverse the linked list and compare the reverse to the original
	2) Iterative approach where we only reverse the first half of the list (Fast runner / slow runner)
	3) Recursive approach Uses a result object
*/
func Palindrome_Attempt1(sll linked_lists.SinglyLinkedList) bool {
	if sll.Length < 2 {
		return false
	}

	odd := sll.Length%2 == 1
	buffer := sll.Length / 2

	currentNode := sll.Head
	firstHalf := make([]int, buffer)
	for ; buffer != 0; buffer-- {
		firstHalf[buffer-1] = currentNode.Data
		currentNode = currentNode.Next
	}

	if odd {
		// if there is an odd length we can progress to the next node
		currentNode = currentNode.Next
	}

	// Compare the values with the remaining nodes
	for _, val := range firstHalf {
		if currentNode.Data != val {
			return false
		}
		currentNode = currentNode.Next
	}


	return true
}
