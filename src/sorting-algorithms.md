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

#### Sorting Test

[Sorting Test](https://github.com/jbelmont/algorithms-workshop/blob/master/sortingAlgorithms/sorting/sorting_test.go)

## Bubble Sort Algorithm

[Bubble Sort (Wikipedia Definition)](https://en.wikipedia.org/wiki/Bubble_sort)

> Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted. The algorithm, which is a comparison sort, is named for the way smaller or larger elements "bubble" to the top of the list.

#### Pseudocode Implementation

[Pseudocode (Wikipedia)](https://en.wikipedia.org/wiki/Bubble_sort#Pseudocode_implementation)

```
procedure bubbleSort(A : list of sortable items)
    n := length(A)
    repeat
        swapped := false
        for i := 1 to n-1 inclusive do
            /* if this pair is out of order */
            if A[i-1] > A[i] then
                /* swap them and remember something changed */
                swap(A[i-1], A[i])
                swapped := true
            end if
        end for
    until not swapped
end procedure
```

#### Bubble Sort Implementation

```go
package bubblesort

func bubblesort(numbers []int, numOfElements int) int {
	var swapTotal int

	for i := 0; i < numOfElements; i++ {

		var numberOfSwaps int

		for j := 0; j < numOfElements-1; j++ {
			if numbers[j] > numbers[j+1] {
				// Use go way to swap numbers instead of longer version in other languages.
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				numberOfSwaps += 1
			}
		}

		swapTotal += numberOfSwaps

		// If no elements were swapped during a traversal, array is sorted
		if numberOfSwaps == 0 {
			break
		}
	}

	return swapTotal
}
```

#### Bubble Sort Test

[Test](https://github.com/jbelmont/algorithms-workshop/blob/master/sortingAlgorithms/bubblesort/bubblesort_test.go)

## Quick Sort Algorithm

[Quick Sort (Wikipedia Definition)](https://en.wikipedia.org/wiki/Quicksort)

> Quicksort (sometimes called partition-exchange sort) is an efficient sorting algorithm. Developed by British computer scientist Tony Hoare in 1959 and published in 1961, it is still a commonly used algorithm for sorting. When implemented well, it can be about two or three times faster than its main competitors, merge sort and heapsort.

#### Lomuto Partitiion Scheme

[Lomuto partition scheme](https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme)

> Lomuto partition scheme
This scheme is attributed to Nico Lomuto and popularized by Bentley in his book Programming Pearls and Cormen et al. in their book Introduction to Algorithms. This scheme chooses a pivot that is typically the last element in the array. The algorithm maintains index i as it scans the array using another index j such that the elements at lo through i-1 (inclusive) are less than the pivot, and the elements at i through j (inclusive) are equal to or greater than the pivot. As this scheme is more compact and easy to understand, it is frequently used in introductory material, although it is less efficient than Hoare's original scheme e.g., when all elements are equal. This scheme degrades to \\( \mathcal{O}(n^2) \\) when the array is already in order.There have been various variants proposed to boost performance including various ways to select pivot, deal with equal elements, use other sorting algorithms such as Insertion sort for small arrays and so on. In pseudocode, a quicksort that sorts elements at lo through hi (inclusive) of an array A can be expressed as:[1

###### Lomuto Pseudocode

```
algorithm quicksort(A, lo, hi) is
    if lo < hi then
        p := partition(A, lo, hi)
        quicksort(A, lo, p - 1)
        quicksort(A, p + 1, hi)

algorithm partition(A, lo, hi) is
    pivot := A[hi]
    i := lo
    for j := lo to hi do
        if A[j] < pivot then
            swap A[i] with A[j]
            i := i + 1
    swap A[i] with A[hi]
    return i
```

#### QuickSort Lomuto Implementation

```go
package quicksort

func partition(numbers []int, low, high int) int {
	i := low
	pivot := numbers[high]
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			// Use go way to swap numbers instead of longer version in other languages.
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i += 1
		}
	}
	numbers[i], numbers[high] = numbers[high], numbers[i]
	return i
}

func quicksort(numbers []int, low, high int) {
	if low < high {
		p := partition(numbers, low, high)
		quicksort(numbers, low, p-1)
		quicksort(numbers, p+1, high)
	}
}
```

#### QuickSort Lomuto Test Run

[Test](https://github.com/jbelmont/algorithms-workshop/blob/master/sortingAlgorithms/quicksort/quicksort_test.go)

## Selection Sort Algorithm

[Selection Sort (Wikipedia Definition)](https://en.wikipedia.org/wiki/Selection_sort)

> In computer science, selection sort is an in-place comparison sorting algorithm. It has an \\( \mathcal{O}(n^2) \\) time complexity, which makes it inefficient on large lists, and generally performs worse than the similar insertion sort. Selection sort is noted for its simplicity and has performance advantages over more complicated algorithms in certain situations, particularly where auxiliary memory is limited.

#### Selection Sort Implementation

```go
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
```

#### Selection Sort Test Run

[Test](https://github.com/jbelmont/algorithms-workshop/blob/master/sortingAlgorithms/selectionsort/selectionsort_test.go)