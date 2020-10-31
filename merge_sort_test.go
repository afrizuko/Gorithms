package gorithms

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	expected := []int{2, 3, 4, 7, 8, 11, 56}
	sorted := MergeSort([]int{3, 4, 2, 56, 7, 8, 11})

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected a sorted array %v but got %v", expected, sorted)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 5; i <= b.N; i++ {
		MergeSort(generateRandomSlice(i))
	}
}

func generateRandomSlice(sliceLength int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(sliceLength)
}
