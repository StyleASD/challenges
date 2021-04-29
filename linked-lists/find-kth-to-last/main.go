package main

import (
	list "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	My implementation of linked list has the length so this becomes an O(N) time problem with no additional space. After
	reviewing the solutions you probably wouldn't have the length because it makes solving this problem trivial.

	After reviewing the answer I misunderstood the question and am returning the wrong values I have altered the code below
	to handle the new information.

	There are several approaches to implementing this without the length. Some use recursion and others return different
	values or print out the node once it is found.

	- Recursive 1 = Prints "kth to last node is head.data" and returns the integer index where the node was found
	- Recursive 2 = where we return the pointer to the node
	- Recursive 3 = We use a wrapper class to iterate over the nodes
	- Iterative 1 = This returns the node without the length


*/
func FindKthToLast_Attempt1(k int, sll list.SinglyLinkedList) int {
	// If there isn't a head
	// or if there isn't k values in the Linked list we can exit
	if sll.Head == nil || sll.Length < k {
		return 0
	}

	currNode := sll.Head
	index := sll.Length - k
	counter := 0
	for counter < index {
		currNode = currNode.Next
		counter++
	}

	return currNode.Data
}
