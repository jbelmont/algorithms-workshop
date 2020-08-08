# Algorithms Workshop

A workshop on Data Structures and Algorithms

## Sections

* [Definition of Algorithm](#definition-of-algorithms)
    * [History of Algorithms](#history-of-algorithms)
    * [Types of Algorithms](#types-of-algorithms)
    * [Types of Data Structures](#types-of-data-structures)
* [Pseudocode](./src/docs/pseudocode.md)
* [Big O Notation](./src/docs/big-o-notation.md)
* [Analysis of Algorithms](./src/docs/analysis-of-algorithms.md)
* [The Halting Problem](./src/docs/the-halting-problem.md)
* [Math for Algorithms](./src/docs/math-for-algorithms.md)
* [Data Structures](./src/docs/data-structures.md)
* [Sorting Algorithms](./src/docs/sorting-algorithms.md)
* [Searching Algorithms](./src/docs/searching-algorithms.md)
* [Graph Algorithms](./src/docs/graph-algorithms.md)
* [String Algorithms](./src/docs/string-algorithms.md)
* [Linear Programming](./src/docs/linear-programming.md)
* [Dynamic Programming](./src/docs/dynamic-programming.md)

## Definition of Algorithms

[Algorithm Definition Wikipedia](https://en.wikipedia.org/wiki/Algorithm)

> In mathematics and computer science, an algorithm (/ˈælɡərɪðəm/ (About this soundlisten)) is a finite sequence of well-defined, computer-implementable instructions, typically to solve a class of problems or to perform a computation. Algorithms are always unambiguous and are used as specifications for performing calculations, data processing, automated reasoning, and other tasks.

Think of algorithms like a recipe in a cookbook. You follow the instructions in the recipe to make your dish.

## History of Algorithms

[Timeline of Algorithms via Wikipedia](https://en.wikipedia.org/wiki/Timeline_of_algorithms)

According to Wikipedia the earliest known algorithm were developed by the Egyptians between 1700 - 2000 BC.

This algorithm involved multiplying 2 numbers together.

In 1600 BC, the Babylonians developed the earliest known algorithm on factorization.

Usually in the university you are introduced to Euclid's Algorithm which was created by Euclid in 300 BC.

Please read the Algorithm Timeline I linked above for a more complete timeline of Algorithms.

## Types of Algorithms

[List of Algorithms](https://en.wikipedia.org/wiki/List_of_algorithms)

Types of Algorithms:

* Simple Recursive Algorithms

* Backtracking Algorithms

* Divide and conquer algorithms

* Dynamic programming algorithms

* Greedy algorithms

* Branch and bound algorithms

* Brute force algorithms

* Randomized algorithms

#### Definition of Recursive Algorithm

[Recursion (Wikipedia)](https://en.wikipedia.org/wiki/Recursion_%28computer_science%29)

> In computer science, recursion is a method of solving a problem where the solution depends on solutions to smaller instances of the same problem. Such problems can generally be solved by iteration, but this needs to identify and index the smaller instances at programming time. Recursion solves such recursive problems by using functions that call themselves from within their own code. The approach can be applied to many types of problems, and recursion is one of the central ideas of computer science.

###### Example of Recursion

```go
func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}
```

Notice here that the fibonacci function is calling itself and had a base case to make sure that the program ends.

#### Backtracking Algorithms

[Backtracking](https://en.wikipedia.org/wiki/Backtracking)

> Backtracking is a general algorithm for finding all (or some) solutions to some computational problems, notably constraint satisfaction problems, that incrementally builds candidates to the solutions, and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly be completed to a valid solution.

Please read more about the [Eight Queens Puzzle](https://en.wikipedia.org/wiki/Eight_queens_puzzle) for a Backtracking Problem.

#### Divide and Conquer Algorithms

[Divide and Conquer Algorithms](https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm)

> In computer science, divide and conquer is an algorithm design paradigm based on multi-branched recursion. A divide-and-conquer algorithm works by recursively breaking down a problem into two or more sub-problems of the same or related type, until these become simple enough to be solved directly. The solutions to the sub-problems are then combined to give a solution to the original problem.

We will look later on in the workshop for specific *Divide and Conquer Algorithms*

#### Dynamic programming algorithms

[Dynamic Programming](https://en.wikipedia.org/wiki/Dynamic_programming)

> Dynamic Programming (DP) is an algorithmic technique for solving an optimization problem by breaking it down into simpler subproblems and utilizing the fact that the optimal solution to the overall problem depends upon the optimal solution to its subproblems.

An example of Dynamic Programming Algorithm is a Dijkstra's Shortest Path Algorithm:

###### Pseudocode Example

```
function Dijkstra(Graph, source):
    create vertex set Q

    for each vertex v in Graph:             
        dist[v] ← INFINITY                  
        prev[v] ← UNDEFINED                 
        add v to Q                      
    dist[source] ←                        

    while Q is not empty:
        u ← vertex in Q with min dist[u]    
                                            
        remove u from Q 
        
        for each neighbor v of u:           // only v that are still in Q
            alt ← dist[u] + length(u, v)
            if alt < dist[v]:               
                dist[v] ← alt 
                prev[v] ← u 

    return dist[], prev[]
```

#### Greedy Algorithm

[Greedy Algorighm Wikipdia](https://en.wikipedia.org/wiki/Greedy_algorithm)

> A greedy algorithm is any algorithm that follows the problem-solving heuristic of making the locally optimal choice at each stage. In many problems, a greedy strategy does not usually produce an optimal solution, but nonetheless a greedy heuristic may yield locally optimal solutions that approximate a globally optimal solution in a reasonable amount of time.

###### Greedy Algorithm Example

*Prim's Algorithm*

Read more about [Prim's Algorithm at Wikipedia](https://en.wikipedia.org/wiki/Prim%27s_algorithm)

#### Branch and Bound Algorithms

[Branch and Bound](https://en.wikipedia.org/wiki/Branch_and_bound)

> Branch and bound (BB, B&B, or BnB) is an algorithm design paradigm for discrete and combinatorial optimization problems, as well as mathematical optimization. A branch-and-bound algorithm consists of a systematic enumeration of candidate solutions by means of state space search: the set of candidate solutions is thought of as forming a rooted tree with the full set at the root. The algorithm explores branches of this tree, which represent subsets of the solution set. Before enumerating the candidate solutions of a branch, the branch is checked against upper and lower estimated bounds on the optimal solution, and is discarded if it cannot produce a better solution than the best one found so far by the algorithm.

###### Branch and Bound Algorithm Example

The [Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem) is an example where you can use Branch and Bound Algorithm to solve.

#### Brute force Algorithm

[Brute Force Search](https://en.wikipedia.org/wiki/Brute-force_search)

> In computer science, brute-force search or exhaustive search, also known as generate and test, is a very general problem-solving technique and algorithmic paradigm that consists of systematically enumerating all possible candidates for the solution and checking whether each candidate satisfies the problem's statement.

#### Randomized Algorithm

[Randomized Algorithm](https://en.wikipedia.org/wiki/Randomized_algorithm)

> A randomized algorithm is an algorithm that employs a degree of randomness as part of its logic. The algorithm typically uses uniformly random bits as an auxiliary input to guide its behavior, in the hope of achieving good performance in the "average case" over all possible choices of random bits. Formally, the algorithm's performance will be a random variable determined by the random bits; thus either the running time, or the output (or both) are random variables.

###### Randomized Algorithm Example

The [Monte Carlo Algorithm](https://en.wikipedia.org/wiki/Monte_Carlo_algorithm) is an example of a Randomized Algorithm.


## Types of Data Structures

Please checkout the [List of Data Structures Wikipedia](https://en.wikipedia.org/wiki/List_of_data_structures) for a complete list of Data Structures by type.

Common Data Structures:

* [Linked List](https://en.wikipedia.org/wiki/Linked_list)

* [Hash Table](https://en.wikipedia.org/wiki/Hash_table)

* [Binary Tree](https://en.wikipedia.org/wiki/Binary_tree)

* [Array](https://en.wikipedia.org/wiki/Array_data_structure)