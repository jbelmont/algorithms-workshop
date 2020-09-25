# Algorithms Workshop - Graph Algorithms

## Graph Theory

[Graph theory (Wikipedia Definition)](https://en.wikipedia.org/wiki/Graph_theory)

> In mathematics, graph theory is the study of graphs, which are mathematical structures used to model pairwise relations between objects. A graph in this context is made up of vertices (also called nodes or points) which are connected by edges (also called links or lines). A distinction is made between undirected graphs, where edges link two vertices symmetrically, and directed graphs, where edges link two vertices asymmetrically; see Graph (discrete mathematics) for more detailed definitions and for other variations in the types of graph that are commonly considered. Graphs are one of the prime objects of study in discrete mathematics.

## Depth-first search algorithm

[DFS (Wikipedia Definition)](https://en.wikipedia.org/wiki/Depth-first_search)

> Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking.

#### Depth-first search recursive pseudocode

```
procedure DFS(G, v) is
    label v as discovered
    for all directed edges from v to w that are in G.adjacentEdges(v) do
        if vertex w is not labeled as discovered then
            recursively call DFS(G, w)
```

#### Depth-first search simple implementation

```go
package depthfirstsearch

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func (n *Node) Search(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.Search(f)
	f(n)
	n.Right.Search(f)
}

func (n *Node) Insert(v string) {
	if v < n.Value {
		if n.Left == nil {
			n.Left = &Node{
				Value: v,
			}
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Value: v,
			}
		} else {
			n.Right.Insert(v)
		}
	}
}
```

#### Depth-first search Test Run

[Depth-first Search Test](https://github.com/jbelmont/algorithms-workshop/blob/master/graphAlgorithms/depthfirstsearch/depthfirstsearch_test.go)

## Social Graph

[Social Graph (Wikipedia Definition)](https://en.wikipedia.org/wiki/Social_graph)

> The social graph is a graph that represents social relations between entities. In short, it is a model or representation of a social network, where the word graph has been taken from graph theory. The social graph has been referred to as "the global mapping of everybody and how they're related".

#### Social Graph Sample Implementation

```go
package networkgraph

type Graph struct {
	Size  int
	Links [][]Link
}

type Link struct {
	VertexA int
	VertexB int
	Weight  int
}

func (g Graph) New(num int) Graph {
	g.Size = num
	g.Links = make([][]Link, num)
	return g
}

func (g *Graph) Add(vertexA int, vertexB int, weight int) {
	g.Links[vertexA] = append(
		g.Links[vertexA],
		Link{
			VertexA: vertexA,
			VertexB: vertexB,
			Weight:  weight,
		},
	)
}
```

#### Social Graph Test Run

[Social Graph Test](https://github.com/jbelmont/algorithms-workshop/blob/master/graphAlgorithms/networkgraph/networkgraph_test.go)