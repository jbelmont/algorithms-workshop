package depthfirstsearch

import (
	"fmt"
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
	node := &Node{}
	node.Insert("c")
	node.Insert("f")
	node.Insert("k")
	returnVal := func(node *Node) {
		fmt.Println(node.Value)
	}
	node.Search(returnVal)
}
