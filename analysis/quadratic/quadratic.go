package quadratic

import "fmt"

func quadraticComplexity(iteration int) int {
	var computation int
	for i := 1; i <= iteration; i++ {
		for j := 1; j <= iteration; j++ {
			computation = i * j
			fmt.Println(computation)
		}
	}
	return computation
}
