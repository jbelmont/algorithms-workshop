package bubblesort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	numbers := []int{
		7, 2, 1, 6, 8, 5, 3, 4,
	}
	bubblesort(numbers, len(numbers))
	expectedNumbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	for index, number := range numbers {
		expectedNumber := expectedNumbers[index]
		if number != expectedNumber {
			t.Errorf("Expected: %d, Actual: %d", expectedNumber, number)
		}
	}
}
