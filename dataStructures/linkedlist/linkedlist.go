package linkedlist

type LinkedList struct {
	Head *Node
}

type Node struct {
	Next     *Node
	Property int
}

func (l *LinkedList) Prepend(property int) {
	node := &Node{}
	node.Property = property
	node.Next = nil

	if l.Head != nil {
		node.Next = l.Head
	}
	l.Head = node
}

func (l *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = l.Head; node != nil; node = node.Next {
		if node.Next == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (l *LinkedList) Append(property int) {
	node := &Node{}
	node.Property = property
	node.Next = nil

	var lastNode *Node
	lastNode = l.LastNode()

	if lastNode != nil {
		lastNode.Next = node
	}
}

func (l *LinkedList) FindNode(property int) *Node {
	var node *Node
	var foundNode *Node

	for node = l.Head; node != nil; node = node.Next {
		if node.Property == property {
			foundNode = node
			break
		}
	}
	return foundNode
}

func (l *LinkedList) InsertAfter(nodeInsertProperty int, property int) {
	node := &Node{}
	node.Property = property
	node.Next = nil

	var insertNode *Node

	insertNode = l.FindNode(nodeInsertProperty)
	if insertNode != nil {
		node.Next = insertNode.Next
		insertNode.Next = node
	}
}
