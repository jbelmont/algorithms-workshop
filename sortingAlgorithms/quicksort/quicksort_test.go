package quicksort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	numbers := []int{
		7, 2, 1, 6, 8, 5, 3, 4,
	}
	low := 0
	high := len(numbers) - 1
	quicksort(numbers, low, high)
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
