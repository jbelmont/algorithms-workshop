package networkgraph

import (
	"fmt"
	"testing"
)

func TestNetworkGraph(t *testing.T) {
	var networkGraph Graph
	networkGraph = networkGraph.New(5)
	networkGraph.Add(0, 2, 1)
	networkGraph.Add(0, 1, 2)
	networkGraph.Add(1, 5, 2)
	networkGraph.Add(1, 7, 3)
	for _, link := range networkGraph.Links {
		for _, l := range link {
			fmt.Println(l.VertexA, l.VertexB, l.Weight)
		}
	}
}
