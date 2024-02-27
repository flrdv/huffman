package main

import (
	"unsafe"
)

// Symbol represents an encoded character
type Symbol struct {
	Value uint32
	Bits  uint8
}

type Queue struct {
	values []Symbol
	index  int
}

func NewQueue() *Queue {
	return &Queue{
		values: make([]Symbol, 1),
		index:  1,
	}
}

func (q *Queue) Push(symbols ...Symbol) {
	q.values = append(q.values, symbols...)
}

func (q *Queue) Pop() Symbol {
	s := q.values[q.index]
	q.index++

	return s
}

func (q *Queue) Return(symbol Symbol) {
	q.index--
	q.values[q.index] = symbol
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) <= q.index
}

func asSymbols(str string) (symbols []Symbol) {
	leaves := Tree(str).Leaves()
	for i := 0; i < len(str); i++ {
		symbols = append(symbols, char2symbol(leaves, str[i]))
	}

	return symbols
}

func char2symbol(leaves []Leaf, char byte) Symbol {
	for _, leaf := range leaves {
		if leaf.Char == char {
			return Symbol{
				Value: leaf.Value,
				Bits:  leaf.Bits,
			}
		}
	}

	panic("char is not found in leaves")
}

func Compress(str string) ([]uint32, int) {
	queue := NewQueue()
	queue.Push(asSymbols(str)...)

	var (
		output []uint32
		batch  uint32
		bits   uint8
	)

	const bitsPerByte = 8
	const bitsPerBatch = uint8(unsafe.Sizeof(batch) * bitsPerByte)

	for !queue.IsEmpty() {
		symbol := queue.Pop()
		if symbol.Bits+bits <= bitsPerBatch {
			batch = (batch << symbol.Bits) | symbol.Value
			bits += symbol.Bits
			continue
		}

		bitsLeft := bitsPerBatch - bits
		batch = (batch << bitsLeft) | (symbol.Value >> (symbol.Bits - bitsLeft))
		leftover := symbol.Value & (0xffff >> (symbol.Bits - bitsLeft))
		queue.Return(Symbol{leftover, symbol.Bits - bitsLeft})
		output = append(output, batch)
		bits = 0
	}

	totalBits := len(output)*int(bitsPerBatch) + int(bits)

	if bits > 0 {
		output = append(output, batch)
	}

	return output, totalBits
}
