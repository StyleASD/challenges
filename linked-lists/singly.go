package main

type SinglyLinkedList struct {
	Head *SingleNode
	Tail *SingleNode
}

type SingleNode struct {
	Data int
	Next *SingleNode
}

func (sll *SinglyLinkedList) AppendNode(data int) {
	var end = &SingleNode{Data: data, Next: nil}
	sll.Tail.Next = end // Attach the Tail to the last element
	sll.Tail = end      // Set Tail to be the new element

}

func (sll *SinglyLinkedList) DeleteNode(d int) {
	currentNode := sll.Head
	if currentNode.Data == d {
		sll.Head = currentNode.Next // Moved To next if this is the element we want to remove
	}

	for currentNode.Next != nil {
		if currentNode.Next.Data == d {
			currentNode.Next = currentNode.Next.Next
		}

		currentNode = currentNode.Next
	}
}
