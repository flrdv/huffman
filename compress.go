package main

import (
	"strconv"
	"unsafe"
)

// Symbol represents an encoded character
type Symbol struct {
	Value uint32
	Bits  uint8
}

func (s Symbol) String() string {
	return pad(strconv.FormatUint(uint64(s.Value), 2), int(s.Bits))
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

func asSymbols(leaves []Leaf, str string) (symbols []Symbol) {
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

func etx(leaves []Leaf) Symbol {
	for _, leaf := range leaves {
		if leaf.ETX {
			return Symbol{
				Value: leaf.Value,
				Bits:  leaf.Bits,
			}
		}
	}

	return Symbol{
		Value: 0,
		Bits:  0,
	}
}

func Compress(str string) ([]uint32, int) {
	leaves := Tree(str).Leaves()
	queue := NewQueue()
	queue.Push(asSymbols(leaves, str)...)
	queue.Push(etx(leaves))

	var (
		output []uint32
		batch  uint32
		bits   uint8
	)

	// assume we're running under conventional machine
	const bitsPerByte = 8
	const bitsPerBatch = uint8(unsafe.Sizeof(batch) * bitsPerByte)

	for !queue.IsEmpty() {
		symbol := queue.Pop()
		if symbol.Bits+bits < bitsPerBatch {
			batch = batch | (symbol.Value << (bitsPerBatch - bits - symbol.Bits))
			//fmt.Println("nrm", format(batch, bitsPerBatch), format(symbol.Value, symbol.Bits), symbol.Bits)
			bits += symbol.Bits
			continue
		}

		bitsLeft := bitsPerBatch - bits
		leftoverBits := symbol.Bits - bitsLeft

		batch = batch | (symbol.Value >> (leftoverBits))
		leftover := symbol.Value << (bitsPerBatch - leftoverBits) >> (bitsPerBatch - leftoverBits)
		queue.Return(Symbol{leftover, leftoverBits})
		output = append(output, batch)
		batch = 0
		bits = 0
	}

	totalBits := len(output)*int(bitsPerBatch) + int(bits)

	if bits > 0 {
		output = append(output, batch)
	}

	return output, totalBits
}

func format(u uint32, bitsize uint8) string {
	return pad(strconv.FormatUint(uint64(u), 2), int(bitsize))
}
