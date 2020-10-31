package gorithms

import (
	"reflect"
	"testing"
)

func TestStackOps(t *testing.T) {
	stack := Stack{}

	if !stack.IsEmpty() {
		t.Errorf("Expected an empty stack but stack was not empty")
	}
	if stack.Size() != 0 {
		t.Errorf("Expected an empty size stack but stack was not empty")
	}

	stack.Push(12)
	stack.Push("John")
	if stack.Size() != 2 {
		t.Errorf("Expected a stack with %d size but got %d", 2, stack.Size())
	}

	peeked := stack.Peek()
	if !reflect.DeepEqual(peeked, "John") {
		t.Errorf("Expected John but found %v", peeked)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected a stack with %d size but got %d", 2, stack.Size())
	}

	result := stack.Pop()
	if stack.Size() != 1 {
		t.Errorf("Expected a stack with %d size but got %d", 1, stack.Size())
	}

	if !reflect.DeepEqual(result, "John") {
		t.Errorf("Expected John but got %v", result)
	}
}
