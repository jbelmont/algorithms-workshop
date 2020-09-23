package binarysearch

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	numbers := []int{
		3, 7, 15, 21, 9, 58, 79, 81,
	}

	number := 79

	predicate := func(i int) bool {
		return numbers[i] >= number
	}
	expectedIndex := 6
	if findIndex := sort.Search(len(numbers), predicate); findIndex != expectedIndex {
		t.Errorf("Expected %d, Received %d\n", expectedIndex, findIndex)
	}
}
