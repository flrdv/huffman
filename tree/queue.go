package tree

type priorityQueue struct {
	nodes []weightedNode
}

func newPriorityQueue() *priorityQueue {
	return new(priorityQueue)
}

func (p *priorityQueue) Push(nodes ...weightedNode) {
	p.nodes = append(p.nodes, nodes...)
}

func (p *priorityQueue) Pop() (n weightedNode) {
	if p.Len() == 0 {
		return n
	}

	n = p.nodes[0]
	index := 0

	for i, node := range p.nodes {
		if node.Frequency < n.Frequency {
			n = node
			index = i
		}
	}

	p.delete(index)

	return n
}

func (p *priorityQueue) delete(i int) {
	copy(p.nodes[i:], p.nodes[i+1:])
	p.nodes = p.nodes[:len(p.nodes)-1]
}

func (p *priorityQueue) Len() int {
	return len(p.nodes)
}
