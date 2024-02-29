package huffman

import (
	"fmt"
	"github.com/flrdv/huffman"
	"github.com/flrdv/huffman/internal/strext"
	"github.com/flrdv/huffman/tree"
	"strconv"
)

func main() {
	const str = "beep boop beer!"

	for _, leaf := range tree.New(str).Leaves() {
		fmt.Println(leaf)
	}

	compressed, bits := huffman.Compress(str)
	for _, u32 := range compressed {
		fmt.Print(strext.PadBinary(strconv.FormatUint(uint64(u32), 2), 32) + " ")
	}
	fmt.Println()
	fmt.Println("total bits:", bits)

	decompressed, err := huffman.Decompress(tree.New(str), compressed)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(strconv.Quote(decompressed))
}
