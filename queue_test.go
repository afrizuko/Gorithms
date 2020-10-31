package gorithms

import (
	"reflect"
	"testing"
)

func TestQueueOps(t *testing.T) {
	q := Queue{}

	if !q.IsEmpty() {
		t.Errorf("Expected an empty queue but queue was not empty")
	}
	if q.Size() != 0 {
		t.Errorf("Expected an empty size queue but queue was not empty")
	}

	q.Enqueue(12)
	q.Enqueue("John")
	if q.Size() != 2 {
		t.Errorf("Expected a queue with %d size but got %d", 2, q.Size())
	}

	result := q.Dequeue()
	if q.Size() != 1 {
		t.Errorf("Expected a queue with %d size but got %d", 1, q.Size())
	}

	if !reflect.DeepEqual(result, 12) {
		t.Errorf("Expected %v but got %v", 12, result)
	}
}
