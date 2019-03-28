package bubblesort

func bubblesort(numbers []int, numOfElements int) int {
	var swapTotal int

	for i := 0; i < numOfElements; i++ {

		var numberOfSwaps int

		for j := 0; j < numOfElements-1; j++ {
			if numbers[j] > numbers[j+1] {
				// Use go way to swap numbers instead of longer version in other languages.
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				numberOfSwaps += 1
			}
		}

		swapTotal += numberOfSwaps

		// If no elements were swapped during a traversal, array is sorted
		if numberOfSwaps == 0 {
			break
		}
	}

	return swapTotal
}
