package huffman

import (
	"github.com/flrdv/huffman/internal/strext"
	"github.com/flrdv/huffman/tree"
	"strconv"
	"unsafe"
)

// Symbol represents an encoded character
type Symbol struct {
	Value uint32
	Bits  uint8
}

func (s Symbol) String() string {
	return strext.PadBinary(strconv.FormatUint(uint64(s.Value), 2), int(s.Bits))
}

func asSymbols(leaves []tree.Leaf, str string) (symbols []Symbol) {
	for i := 0; i < len(str); i++ {
		symbols = append(symbols, char2symbol(leaves, str[i]))
	}

	return symbols
}

func char2symbol(leaves []tree.Leaf, char byte) Symbol {
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

func etx(leaves []tree.Leaf) Symbol {
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
	leaves := tree.New(str).Leaves()
	queue := newReturnQueue()
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
