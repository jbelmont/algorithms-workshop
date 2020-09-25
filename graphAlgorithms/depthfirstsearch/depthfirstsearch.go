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
