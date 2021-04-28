package linked_lists

type DoublyLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
	Length int
}

type DoubleNode struct {
	Data int
	Next *DoubleNode
	Prev *DoubleNode
}

func NewDoublyLinkedList(data ...int) *DoublyLinkedList {
	list := &DoublyLinkedList{Length: 0}

	for _, datum := range data {
		list.AppendNode(datum)
	}

	return list
}

func (dll *DoublyLinkedList) AppendNode(data int) {
	end := &DoubleNode{Data: data, Prev: dll.Tail, Next: nil}

	dll.Length++;
	if dll.Head == nil {
		dll.Head = end
	} else {
		dll.Tail.Next = end // Attach the Tail to the last element
	}

	dll.Tail = end      // Set Tail to be the new element
}

func (dll *DoublyLinkedList) DeleteNode(data int) {
	currentNode := dll.Head
	if dll.Head == nil {
		return
	}

	dll.Length--;
	if currentNode.Data == data {
		currentNode.Prev = nil      // This is the new head so we remove prev
		dll.Head = currentNode.Next // Moved To next if this is the element we want to remove
		return                      // Exit cause te work is done
	}

	for currentNode.Next != nil {
		if currentNode.Next.Data == data {
			next := currentNode.Next.Next

			next.Prev = currentNode // fix pointer to the prev node
			currentNode.Next = next // fix pointer to the current node
			return                  // Exit cause the work is done
		}

		currentNode = currentNode.Next
	}
}
