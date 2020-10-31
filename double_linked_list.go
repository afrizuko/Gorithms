package gorithms

type DoublyLinkedListNode struct {
	value interface{}
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

func (d *DoublyLinkedListNode) NextNode(node *DoublyLinkedListNode) {
	d.next = node
}
func (d *DoublyLinkedListNode) PrevNode(node *DoublyLinkedListNode) {
	d.prev = node
}
