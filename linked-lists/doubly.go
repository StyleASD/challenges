package main

type DoublyLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

type DoubleNode struct {
	Data int
	Next *DoubleNode
	Prev *DoubleNode
}

func (dll *DoublyLinkedList) AppendNode(data int) {
	end := &DoubleNode{Data: data, Prev: dll.Tail, Next: nil}

	dll.Tail.Next = end // Attach the Tail to the last element
	dll.Tail = end      // Set Tail to be the new element
}

func (dll *DoublyLinkedList) DeleteNode(data int) {
	currentNode := dll.Head
	if currentNode.Data == data {
		dll.Head = currentNode.Next // Moved To next if this is the element we want to remove
		currentNode.Prev = nil      // This is the new head so we remove prev
		return // Exit cause te work is done
	}

	for currentNode.Next != nil {
		if currentNode.Next.Data == data {

			next := currentNode.Next.Next

			currentNode.Next = next // fix pointer to the current node
			next.Prev = currentNode // fix pointer to the prev node
			return // Exit cause the work is done
		}

		currentNode = currentNode.Next
	}
}
