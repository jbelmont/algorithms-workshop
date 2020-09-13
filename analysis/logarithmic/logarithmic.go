package logarithmic

import "fmt"

type Tree struct {
	LNode *Tree
	Val   int
	RNode *Tree
}

func (t *Tree) insert(val int) {
	if t != nil {
		if t.LNode == nil {
			t.LNode = &Tree{
				nil, val, nil,
			}
		} else if t != nil && t.RNode == nil {
			t.RNode = &Tree{
				nil, val, nil,
			}
		} else {
			if t.LNode != nil {
				t.LNode.insert(val)
			} else {
				t.RNode.insert(val)
			}
		}
	} else {
		t = &Tree{
			nil, val, nil,
		}
	}
}

func printTree(t *Tree) {
	if t != nil {
		fmt.Printf("Tree value is %v\n", t.Val)
		fmt.Printf("Left Tree Node ")
		printTree(t.LNode)
		fmt.Printf("Right Tree Node ")
		printTree(t.RNode)
	} else {
		fmt.Printf("Nil\n")
	}
}
