# Algorithms Workshop - Search Algorithms

## Search Algorithm Definition

[Search Algorithm (Wikipedia Definition)](https://en.wikipedia.org/wiki/Search_algorithm)

> In computer science, a search algorithm is any algorithm which solves the search problem, namely, to retrieve information stored within some data structure, or calculated in the search space of a problem domain, either with discrete or continuous values. Specific applications of search algorithms include:

> Problems in combinatorial optimization, such as:
>    * The vehicle routing problem, a form of shortest path problem
>    * The knapsack problem: Given a set of items, each with a weight and a value, determine the number of each item to include in a collection so that the total weight is less than or equal to a given limit and the total value is as large as possible.
>    * The nurse scheduling problem
>    * Problems in constraint satisfaction, such as:
>    * The map coloring problem
>    * Filling in a sudoku or crossword puzzle
>    * In game theory and especially combinatorial game theory, choosing the best move to make next
>    * Finding a combination or password from the whole set of possibilities
>    * Factoring an integer (an important problem in cryptography)
>    * Optimizing an industrial process, such as a chemical reaction, by changing the parameters of the process
>    * Retrieving a record from a database
>    * Finding the maximum or minimum value in a list or array
>    * Checking to see if a given value is present in a set of values

## Linear Search Algorithm

[Linear Search (Wikipedia Definition)](https://en.wikipedia.org/wiki/Linear_search)

> In computer science, a linear search or sequential search is a method for finding an element within a list. It sequentially checks each element of the list until a match is found or the whole list has been searched.

#### Linear Search Basic Algorithm

[Basic Algorithm](https://en.wikipedia.org/wiki/Linear_search#Basic_algorithm)

![Linear Search](./images/linear_search_algorithm.png)

#### Linear Search Implementation

```go
package linearsearch

func search(elements []int, searchElem int) int {
	for index, element := range elements {
		if element == searchElem {
			return index
		}
	}
	return -1
}
```

#### Linear Search Test

[Linear Search Test](https://github.com/jbelmont/algorithms-workshop/blob/master/searchingAlgorithms/linearsearch/linearsearch_test.go)