package main

import (
	"fmt"
	"strconv"
)

func main() {
	const str = "beep boop beer!"

	for _, leaf := range Tree(str).Leaves() {
		fmt.Println(leaf)
	}

	compressed, bits := Compress(str)
	for _, u32 := range compressed {
		fmt.Print(pad(strconv.FormatUint(uint64(u32), 2), 32) + " ")
	}
	fmt.Println()
	fmt.Println("total bits:", bits)

	decompressed, err := Decompress(Tree(str), compressed)
	fmt.Println("err", err, "str", strconv.Quote(decompressed), len(str), len(decompressed))
}
