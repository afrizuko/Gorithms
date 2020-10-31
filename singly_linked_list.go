package gorithms

type SinglyLinkedList struct {
	val      interface{}
	nextnode *SinglyLinkedList
}

func (s SinglyLinkedList) Size() int {
	return 0
}

func (s *SinglyLinkedList) NextNode(nextNode *SinglyLinkedList) {
	s.nextnode = nextNode
}
