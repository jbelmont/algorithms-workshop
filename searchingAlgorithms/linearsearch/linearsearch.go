package linearsearch

func search(elements []int, searchElem int) int {
	for index, element := range elements {
		if element == searchElem {
			return index
		}
	}
	return -1
}
