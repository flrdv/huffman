package main

type PriorityQueue struct {
	nodes []WeightedNode
}

func NewPriorityQueue() *PriorityQueue {
	return new(PriorityQueue)
}

func (p *PriorityQueue) Push(nodes ...WeightedNode) {
	p.nodes = append(p.nodes, nodes...)
}

func (p *PriorityQueue) Pop() (n WeightedNode) {
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

func (p *PriorityQueue) delete(i int) {
	copy(p.nodes[i:], p.nodes[i+1:])
	p.nodes = p.nodes[:len(p.nodes)-1]
}

func (p *PriorityQueue) Len() int {
	return len(p.nodes)
}
