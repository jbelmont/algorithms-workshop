# Algorithms Workshop - Sorting Algorithms

## Sorting Algorithm Definition

[Sorting Algorithm (Wikipedia Definition)](https://en.wikipedia.org/wiki/Sorting_algorithm)

> In computer science, a sorting algorithm is an algorithm that puts elements of a list in a certain order. The most frequently used orders are numerical order and lexicographical order. Efficient sorting is important for optimizing the efficiency of other algorithms (such as search and merge algorithms) that require input data to be in sorted lists. Sorting is also often useful for canonicalizing data and for producing human-readable output. More formally, the output of any sorting algorithm must satisfy two conditions:

>  The output is in nondecreasing order (each element is no smaller than the previous element according to the desired total order);

>  The output is a permutation (a reordering, yet retaining all of the original elements) of the input.

## Sorting in Golang

In Golang you can use the [Sort package](https://pkg.go.dev/sort) to sort slices

You define the following methods:

* Len
* Swap
* Less

#### Sorting Code in Action

```go
package sorting

import (
	"fmt"
)

type Soldier struct {
	Name           string
	Rank           string
	Age            int
	YearsOfService int
}

func (s Soldier) ToString() string {
	return fmt.Sprintf(
		"Name: %s\nRank: %s\nAge: %d\nYearsOfService: %d\n",
		s.Name,
		s.Rank,
		s.Age,
		s.YearsOfService,
	)
}

type SortByYearsOfService []Soldier

func (s SortByYearsOfService) Len() int {
	return len(s)
}

func (s SortByYearsOfService) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByYearsOfService) Less(i int, j int) bool {
	return s[i].YearsOfService < s[j].YearsOfService
}
```

#### Sorting Test Run

[Sorting Test Run](https://github.com/jbelmont/algorithms-workshop/blob/master/sortingAlgorithms/sorting/sorting_test.go)