package main

import (
	linked_lists "programming-problems/linked-lists"
)

func main() {

}

func LoopDetection_Attempt1(list1 linked_lists.SinglyLinkedList) *linked_lists.SingleNode {
	if list1.Head == nil {
		return nil
	}

	currentNode := list1.Head
	hash := map[*linked_lists.SingleNode]struct{}{}

	for currentNode != nil {
		if _, ok := hash[currentNode]; ok {
			return currentNode
		}

		hash[currentNode] = struct{}{}
		currentNode = currentNode.Next
	}

	return nil
}
