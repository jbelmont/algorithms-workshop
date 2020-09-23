package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	numbers := []int{
		1, 5, 7, 9, 8, 17,
	}
	if search(numbers, 8) != 4 {
		t.Error("Index should equal 4")
	}

	if search(numbers, 100) != -1 {
		t.Error("Should give a positive index")
	}
}
