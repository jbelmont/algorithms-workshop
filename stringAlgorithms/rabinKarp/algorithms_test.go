package stringAlgorithms

import (
	"fmt"
	"testing"
)

func TestRabinKarp(t *testing.T) {
	rabinKarp := NewRabinKarp("needle")
	byteString := []byte("needle in a Haystack")

	var num int
	for _, c := range byteString {
		if receiveChar := rabinKarp.SearchForNextCharacter(c); receiveChar {
			num++
		}
	}

	fmt.Println(num)
}
