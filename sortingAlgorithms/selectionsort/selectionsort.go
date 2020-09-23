package selectionsort

func selectionSort(elements []int) {
	count := len(elements) - 1
	for i := 0; i < count; i++ {
		minimum := i
		var j int
		for j = i + 1; j <= count; j++ {
			if elements[j] < elements[minimum] {
				minimum = j
			}
		}
		// Swap out elements
		elements[i], elements[minimum] = elements[minimum], elements[i]
	}
}
