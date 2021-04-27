package main

type SinglyLinkedList struct {
	Date int
	Next *SinglyLinkedList
}

func (head *SinglyLinkedList) AppendNode(data int) {
	var end = &SinglyLinkedList{Date: data, Next: nil}
	n := head

	for head.Next != nil {
		n = n.Next
	}

	n.Next = end

}

func (head *SinglyLinkedList) DeleteNode(d int) *SinglyLinkedList {
	if head == nil {
		return nil
	}

	n := head

	if n.Date == d {
		return head.Next // Moved Head
	}

	for n.Next != nil {
		if n.Next.Date == d {
			n.Next = n.Next.Next
			return head // head didnt change
		}

		n = n.Next
	}

	return head
}