package main

import (
	linked_lists "programming-problems/linked-lists"
)

func main() {

}

/*
	Attempt 1:

	For some reason I found this one difficult. I ended up looking at the answer and copying it below. The idea I had
	was similar to this bout i couldn't figure out how to walk the tail and prepend to the head value.

	Basically the way the below works if the number is bigger it gets moved to the tail, if the number is smaller it gets
	moved to the head.

*/
func Partition_Attempt1(sll *linked_lists.SinglyLinkedList, partition int) {
	currentNode := sll.Head
	head := sll.Head
	tail := sll.Head

	for currentNode != nil {
		next := currentNode.Next // Save the next node so we can pass it to the next loop

		/*
			Check the partition and depending on the value either prepend it to the head or append it to the tail
		*/
		if currentNode.Data < partition {
			currentNode.Next = head
			head = currentNode
		} else {
			tail.Next = currentNode
			tail = currentNode
		}

		currentNode = next
	}
	// Tidying up after we are done by setting the tail value to be nil and passing head back to the sll
	tail.Next = nil
	sll.Head = head
}
