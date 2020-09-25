# Algorithms Workshop - Dynamic Programming

## Dynamic Programming Definition

[Dynamic Programming (Wikipedia Definition)](https://en.wikipedia.org/wiki/Dynamic_programming)

> Dynamic programming is both a mathematical optimization method and a computer programming method. The method was developed by Richard Bellman in the 1950s and has found applications in numerous fields, from aerospace engineering to economics.

> In both contexts it refers to simplifying a complicated problem by breaking it down into simpler sub-problems in a recursive manner. While some decision problems cannot be taken apart this way, decisions that span several points in time do often break apart recursively. Likewise, in computer science, if a problem can be solved optimally by breaking it into sub-problems and then recursively finding the optimal solutions to the sub-problems, then it is said to have optimal substructure.

## Dynamic Programming Algorithms and Problems

* [Dijkstra's Algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
* [Fibonacci Numbers and Sequence](https://en.wikipedia.org/wiki/Fibonacci_sequence)
* [Tower of Hanoi](https://en.wikipedia.org/wiki/Tower_of_Hanoi)

#### Dijkstra's Algorithm

> Dijkstra's algorithm (or Dijkstra's Shortest Path First algorithm, SPF algorithm) is an algorithm for finding the shortest paths between nodes in a graph, which may represent, for example, road networks. It was conceived by computer scientist Edsger W. Dijkstra in 1956 and published three years later.

> The algorithm exists in many variants. Dijkstra's original algorithm found the shortest path between two given nodes, but a more common variant fixes a single node as the "source" node and finds shortest paths from the source to all other nodes in the graph, producing a shortest-path tree.

###### Dijkstra Algorithm Pseudocode

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

##### Dijkstra Golang Implementation

Check out this blog post on [Dijkstra Algorithm](https://deployeveryday.com/2019/10/16/dijkstra-algorithm-golang.html)