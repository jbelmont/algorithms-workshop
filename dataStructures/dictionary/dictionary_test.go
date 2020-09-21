package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := &Dictionary{}

	dictionary.Put("name", "John Rambo")
	dictionary.Put("rank", "SFC")
	dictionary.Put("age", "32")

	dictionary.Remove("age")

	if !dictionary.Contains("rank") {
		t.Error("Should contain rank keyword in dictionary")
	}

	if string(dictionary.Find("rank")) != "SFC" {
		t.Error("Should have rank of SFC")
	}

	if dictionary.Count() != 2 {
		t.Error("Should have count of 2")
	}

	keys := dictionary.GetKeys()
	if string(keys[0]) != "name" {
		t.Error("First key should equal \"name\"")
	}

	values := dictionary.GetValues()
	if string(values[0]) != "John Rambo" {
		t.Error("First value in key should equal \"John Rambo\"")
	}
}
