package main

import (
	linked_lists "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	This ones is fairly simple to resolve for the middle because we can just copy the data from the next node and change
	the next pointer.
*/
func DeleteMiddleNode_Attempt1(n *linked_lists.SingleNode) {
	if n == nil || n.Next == nil {
		return // Exit early here because the node can't be changed
	}

	// We can remove the node without having to access to the whole linked list
	// By ensuring
	n.Data = n.Next.Data
	n.Next = n.Next.Next

}
