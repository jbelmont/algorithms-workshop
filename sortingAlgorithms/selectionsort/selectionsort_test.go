package selectionsort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	elements := []int{
		19, 29, 57, 89, 8, 21, 77, 69,
	}

	selectionSort(elements)
	expected := []int{
		8, 19, 21, 29, 57, 69, 77, 89,
	}
	for index, value := range elements {
		if value != expected[index] {
			t.Errorf("Expected %d, Received %d\n", expected[index], value)
		}
	}
}
