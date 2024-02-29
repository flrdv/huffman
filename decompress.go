package huffman

import (
	"fmt"
	"github.com/flrdv/huffman/tree"
)

func Decompress(root *tree.Node, data []uint32) (string, error) {
	var decompressed []byte
	queue := NewBitQueue()
	queue.Push(data...)
	node := root

	for !queue.IsEmpty() {
		bit := queue.Pop()
		if bit == 0 {
			node = node.Left
		} else {
			node = node.Right
		}

		if node.ETX {
			return string(decompressed), nil
		}

		if node.Char != '\x00' {
			decompressed = append(decompressed, node.Char)
			node = root
		}
	}

	return "", fmt.Errorf("no ETX marker")
}
