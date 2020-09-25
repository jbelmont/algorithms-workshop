package brutesearch

import (
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	text := "abracadbra"
	pattern := "d"
	index := search(text, pattern)
	expected := 6
	if index != expected {
		t.Errorf("Expected %d, Received: %d", expected, index)
	}

	// Search Function from Standard Library
	foundIndex := strings.Index(text, pattern)
	if foundIndex != expected {
		t.Errorf("Expected %d, Received: %d", expected, foundIndex)
	}
}
