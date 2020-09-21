package queue

type Queue struct {
	Nodes []*Node
	Size  int
	Head  int
	Tail  int
	Count int
}

type Node struct {
	Value int
	Name  string
}

// New returns a new queue with given size
func (q *Queue) New(size int) *Queue {
	q.Nodes = make([]*Node, size)
	q.Size = size
	return q
}

// Enqueue addes a node to the queue
func (q *Queue) Enqueue(node *Node) {
	if q.Head == q.Tail && q.Count > 0 {
		nodes := make([]*Node, len(q.Nodes)+q.Size)
		copy(nodes, q.Nodes[q.Head:])
		copy(
			nodes[len(q.Nodes)-q.Head:],
			q.Nodes[:q.Head],
		)

		q.Head = 0
		q.Tail = len(q.Nodes)
		q.Nodes = nodes
	}
	q.Nodes[q.Tail] = node
	q.Tail = (q.Tail + 1) % len(q.Nodes)
	q.Count++
}

// Dequeue removes and returns the node in the queue in LIFO order
// Last In First Out
func (q *Queue) Dequeue() *Node {
	if q.Count == 0 {
		return nil
	}
	node := q.Nodes[q.Head]
	q.Head = (q.Head + 1) % len(q.Nodes)
	q.Count--
	return node
}
