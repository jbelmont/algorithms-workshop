package cubic

import (
	"fmt"
)

func cubicComplexity() int {
	const iteration = 10
	var threeDArr [iteration][iteration][iteration]int
	var compute int
	for i := 0; i < iteration; i++ {
		for j := 0; j < iteration; j++ {
			for k := 0; k < iteration; k++ {
				compute = (i + j) * k
				threeDArr[i][j][k] = compute
				fmt.Printf("Value in threeDArr[i][j][k] = %d\n", threeDArr[i][j][k])
			}
		}
	}
	return compute
}
