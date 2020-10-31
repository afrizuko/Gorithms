package gorithms

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	a := DoublyLinkedListNode{value: 3}
	b := DoublyLinkedListNode{value: 3}
	c := DoublyLinkedListNode{value: 3}
	d := DoublyLinkedListNode{value: 3}

	a.NextNode(&b)
	b.NextNode(&c)
	b.PrevNode(&a)
	c.NextNode(&d)
	c.PrevNode(&b)

	if a.next != &b {
		t.Errorf("Expected a %v but got %v", &b, a.next)
	}

	if a.value != 3 {
		t.Errorf("Expected a %v but got %v", &b, a.next)
	}
}
