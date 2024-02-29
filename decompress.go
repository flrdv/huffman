package main

import (
	"unsafe"
)

type BitQueue struct {
	index    int
	values   []uint32
	bitShift uint8
}

func NewBitQueue() *BitQueue {
	return new(BitQueue)
}

func (b *BitQueue) Push(values ...uint32) {
	b.values = append(b.values, values...)
}

func (b *BitQueue) Pop() byte {
	if b.IsEmpty() {
		panic("")
		return 0
	}

	const bitsPerValue = uint8(unsafe.Sizeof(b.values[0])) * 8
	bit := byte(b.values[b.index]>>(bitsPerValue-b.bitShift-1)) & 1

	b.bitShift++
	if b.bitShift > bitsPerValue {
		b.bitShift = 0
		b.index++
	}

	return bit
}

func (b *BitQueue) IsEmpty() bool {
	return b.index >= len(b.values)
}

func Decompress(root *Node, data []uint32) (string, error) {
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

		switch node.Char {
		case '\x00':
		case ETX:
			return string(decompressed), nil
		default:
			decompressed = append(decompressed, node.Char)
			node = root
		}
	}

	//if node != root {
	//	return "", fmt.Errorf("incomplete sequence")
	//}

	return string(decompressed), nil
}
