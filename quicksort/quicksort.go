package quicksort

func partition(numbers []int, low, high int) int {
	i := low
	pivot := numbers[high]
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			// Use go way to swap numbers instead of longer version in other languages.
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i += 1
		}
	}
	numbers[i], numbers[high] = numbers[high], numbers[i]
	return i
}

func quicksort(numbers []int, low, high int) {
	if low < high {
		p := partition(numbers, low, high)
		quicksort(numbers, low, p-1)
		quicksort(numbers, p+1, high)
	}
}
