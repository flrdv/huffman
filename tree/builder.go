package tree

import (
	"fmt"
	"github.com/flrdv/huffman/internal/strext"
	"strconv"
)

type Node struct {
	Left, Right *Node
	Char        byte
	ETX         bool
}

func (n *Node) String() string {
	return fmt.Sprintf("Node(char=%s)", strconv.Quote(string(n.Char)))
}

func New(str string) *Node {
	queue := newPriorityQueue()
	queue.Push(countLetters(str)...)
	queue.Push(newETXNode())

	for queue.Len() > 1 {
		a, b := queue.Pop(), queue.Pop()
		queue.Push(newWeightedNode('\x00', a.Frequency+b.Frequency, a.Node, b.Node))
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
		char, strext.PadBinary(strconv.FormatUint(uint64(l.Value), 2), int(l.Bits)),
	)
}

type weightedNode struct {
	*Node
	Frequency int
}

func (w weightedNode) String() string {
	return fmt.Sprintf("%d(%s)", w.Frequency, w.Node.String())
}

func newWeightedNode(char byte, freq int, left, right *Node) weightedNode {
	return weightedNode{
		Node: &Node{
			Char:  char,
			Left:  left,
			Right: right,
		},
		Frequency: freq,
	}
}

func newETXNode() weightedNode {
	return weightedNode{Node: &Node{
		ETX: true,
	}}
}
