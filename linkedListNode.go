package linkedlist

// A linked list node is a single value in a linked list
type LinkedListNode[contained any] struct {
	Previous *LinkedListNode[contained]
	Next     *LinkedListNode[contained]

	Value contained
}

// adds a node with value as its value after the node, returns this node
func (n *LinkedListNode[contained]) Append(value contained, list *LinkedList[contained]) (appended *LinkedListNode[contained]) {
	appended = new(LinkedListNode[contained])

	appended.Next = n.Next
	appended.Previous = n

	if appended.Next == nil {
		list.Last = appended

		if appended.Next == nil {
			list.First = appended
		}
	}

	n.Next = appended

	return
}

// adds a node with value as its value before the node, returns this node
func (n *LinkedListNode[contained]) Prepend(value contained, list *LinkedList[contained]) (prepended *LinkedListNode[contained]) {
	prepended = new(LinkedListNode[contained])

	prepended.Next = n
	prepended.Previous = n.Previous

	if prepended.Previous == nil {
		list.First = prepended

		if prepended.Next == nil {
			list.Last = prepended
		}
	}

	n.Previous = n

	return
}

//removes the node and joins the values before and after it, returns the value of the deleted node
func (n *LinkedListNode[contained]) Pop() (value contained) {
	n.Next.Previous = n.Previous
	n.Previous.Next = n.Next

	value = n.Value

	return
}
