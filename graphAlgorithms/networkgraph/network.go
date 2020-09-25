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
