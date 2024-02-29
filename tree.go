package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Left, Right *Node
	Char        byte
	ETX         bool
}

func (n *Node) String() string {
	return fmt.Sprintf("Node(char=%s)", strconv.Quote(string(n.Char)))
}

type WeightedNode struct {
	*Node
	Frequency int
}

func (w WeightedNode) String() string {
	return fmt.Sprintf("%d(%s)", w.Frequency, w.Node.String())
}

func NewWeightedNode(char byte, freq int, left, right *Node) WeightedNode {
	return WeightedNode{
		Node: &Node{
			Char:  char,
			Left:  left,
			Right: right,
		},
		Frequency: freq,
	}
}

func NewETXNode() WeightedNode {
	return WeightedNode{Node: &Node{
		ETX: true,
	}}
}

func Tree(str string) *Node {
	queue := NewPriorityQueue()
	queue.Push(countLetters(str)...)
	queue.Push(NewETXNode())

	for queue.Len() > 1 {
		a, b := queue.Pop(), queue.Pop()
		queue.Push(NewWeightedNode('\x00', a.Frequency+b.Frequency, a.Node, b.Node))
	}

	return queue.Pop().Node
}

func (n *Node) Leaves() (leaves []Leaf) {
	return n.leaves(0, 0)
}

func (n *Node) leaves(mask uint32, depth uint8) (leaves []Leaf) {
	if n.Char != '\x00' || n.ETX {
		// the node is marked as a leaf, so it cannot have any children
		return []Leaf{{n.Char, depth, n.ETX, mask}}
	}

	if n.Left != nil {
		leaves = append(leaves, n.Left.leaves(mask<<1|0, depth+1)...)
	}

	if n.Right != nil {
		leaves = append(leaves, n.Right.leaves(mask<<1|1, depth+1)...)
	}

	return leaves
}

type Leaf struct {
	Char  byte
	Bits  uint8
	ETX   bool
	Value uint32
}

func (l Leaf) String() string {
	char := string(l.Char)
	switch {
	case char == " ":
		char = "<SP>"
	case l.ETX:
		char = "<ETX>"
	}

	return fmt.Sprintf(
		"%s: %s",
		char, pad(strconv.FormatUint(uint64(l.Value), 2), int(l.Bits)),
	)
}

func pad(str string, desired int) string {
	if len(str) >= desired {
		return str
	}

	return strings.Repeat("0", desired-len(str)) + str
}

func countLetters(str string) (nodes []WeightedNode) {
	for i := 0; i < len(str); i++ {
		index := nodeIndex(str[i], nodes)
		if index == -1 {
			nodes = append(nodes, NewWeightedNode(str[i], 1, nil, nil))
		} else {
			nodes[index].Frequency++
		}
	}

	return nodes
}

func nodeIndex(char byte, nodes []WeightedNode) int {
	for i, node := range nodes {
		if node.Char == char {
			return i
		}
	}

	return -1
}
