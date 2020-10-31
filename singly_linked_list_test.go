package gorithms

import "testing"

func TestSinglyLinkedList(t *testing.T) {
	a := SinglyLinkedList{val: 12}

	want := 0
	got := a.Size()
	if got != want {
		t.Errorf("Expected a size of %d but got %d", want, got)
	}

	b := SinglyLinkedList{val: 15}
	a.NextNode(&b)

	c := SinglyLinkedList{val: 16}
	b.NextNode(&c)

	if b.nextnode != &c {
		t.Errorf("expected %v but got %v", c, b.nextnode)
	}

	if b.val != 15 {
		t.Errorf("expected %v but got %v", c, b.nextnode)
	}
}
