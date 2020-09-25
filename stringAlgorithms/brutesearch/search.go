package brutesearch

func search(text string, pattern string) int {
	var index int
	for i, ch := range text {
		if string(ch) == pattern {
			index = i
		}
	}
	return index
}
