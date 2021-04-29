package main

import (
	list "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	This may be easier because of my implementation of linked list but I can do a check on the length without the length
	it could still be done in O(N) time by looping by keeping an eye on the currNode and ensuring that it isn't nil.

	//


*/
func FindKthToLast_Attempt1(k int, sll list.SinglyLinkedList) int {
	// If there isn't a head
	// or if there isn't k values in the Linked list we can exit
	if sll.Head == nil || sll.Length < k {
		return 0
	}

	currNode := sll.Head
	counter := 1
	for counter != k {
		currNode = currNode.Next
		counter++
	}

	return currNode.Data
}
