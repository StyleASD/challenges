package main

import (
	linked_lists "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1

	So the request is to check the reference to find out an intersection. I dont really see an algorithm which would work
	which is better then O(A + B) with an O(A) in space complexity.

	Answer Review
	My answer is correct in that we can do this however there is an easier way of doing this because we know that the an
	intersection must end the same we can loop through the nodes and comparing them as we go. To do this we will need to
	compare the lengths and chop of any access nodes.

	With the above this should be possible with an O(A + B) time and O(1) space

	In Future
	I should consider the nature of a linked list and how any intersecting nodes much end the same
*/
func Intersection_Attempt1(list1, list2 linked_lists.SinglyLinkedList) *linked_lists.SingleNode {
	if list1.Head == nil || list2.Head == nil {
		return nil
	}

	hashSet := map[*linked_lists.SingleNode]struct{}{}

	// Loop through both finding an intersection
	currNode := list1.Head
	for currNode != nil {
		hashSet[currNode] = struct{}{}
		currNode = currNode.Next
	}

	// Check For intersections
	currNode = list2.Head
	for currNode != nil {
		if _, ok := hashSet[currNode]; ok {
			return currNode
		}
		currNode = currNode.Next
	}

	return nil
}
