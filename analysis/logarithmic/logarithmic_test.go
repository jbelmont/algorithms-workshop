package logarithmic

import (
	"testing"
)

func TestTreeOps(t *testing.T) {
	tree := &Tree{
		nil,
		1,
		nil,
	}

	printTree(tree)

	tree.insert(5)
	printTree(tree)

	tree.insert(10)
	printTree(tree)
}
