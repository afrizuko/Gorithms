package gorithms

import (
	"testing"
)

func TestFibonacciGenerator(t *testing.T) {

	f := fib()
	firstSet := []int{0, 1, 1, 2, 3, 5}
	for i := 0; i <= 5; i++ {
		result := f()
		if result != firstSet[i] {
			t.Errorf("Expected next number in series to be %d but got %d ", firstSet[i], result)
		}
	}
}

func BenchmarkFibonachi(b *testing.B) {
	f := fib()
	// Function calls are evaluated left-to-right.
	for i := 0; i <= b.N; i++ {
		f()
	}
}
