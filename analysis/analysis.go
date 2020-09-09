package main

import (
	"fmt"
)

func main() {
	var number [20]int

	for n := 0; n < 20; n++ {
		number[n] = n + 100
		fmt.Printf("Number[%d] = %d\n", n, number[n])
	}
}
