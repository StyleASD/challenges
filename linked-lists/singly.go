package linked_lists

type SinglyLinkedList struct {
	Head   *SingleNode
	Tail   *SingleNode
	Length int
}

type SingleNode struct {
	Data int
	Next *SingleNode
}

func NewSinglyLinkedList(data ...int) *SinglyLinkedList {
	list := &SinglyLinkedList{Length: 0}

	for _, datum := range data {
		list.AppendData(datum)
	}

	return list
}

/*
	Note there is a bug in this implementation if we have intersecting linked lists. This will work correctly for the
	list calling the end but the tail will be incorrect for another linked list with intersecting nodes.
 */
func (sll *SinglyLinkedList) AppendNode(end *SingleNode) {
	sll.Length++
	if sll.Head == nil {
		sll.Head = end
	} else {
		sll.Tail.Next = end // Set Tail to be the new element
	}

	sll.Tail = end
}

func (sll *SinglyLinkedList) AppendData(data int) {
	var end = &SingleNode{Data: data, Next: nil}

	sll.Length++
	if sll.Head == nil {
		sll.Head = end
	} else {
		sll.Tail.Next = end // Set Tail to be the new element
	}

	sll.Tail = end
}

func (sll *SinglyLinkedList) DeleteNode(d int) {
	currentNode := sll.Head

	if sll.Head == nil {
		return
	}

	sll.Length--
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
