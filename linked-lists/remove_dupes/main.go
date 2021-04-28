package main

import (
	list "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	The question mentions using a buffer so I think this can probably be done in O(N) time. If we store whether we have
	seen a result previously this will allow us to remove values on the fly form the linked list.

	If we are not allowed to used a hash table as a store then we are limited to an O(N^2) because we will have to have
	two loops to check all values against every other value in the linked list.


*/
func RemoveDupes_FirstAttempt(linkedList *list.SinglyLinkedList) {
	if linkedList.Length < 1 {
		return // There is no duplicates if the length of the linked list is less than 1
	}

	var hash = make(map[int]struct{})

	prevNode := linkedList.Head
	hash[prevNode.Data] = struct{}{}
	for prevNode.Next != nil {
		currNode := prevNode.Next

		if _, ok := hash[currNode.Data]; ok {
			// Remove
			prevNode.Next = currNode.Next
			linkedList.Length--
		} else {
			prevNode = currNode
			hash[prevNode.Data] = struct{}{}
		}
	}

}
