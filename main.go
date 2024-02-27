package main

import (
	"fmt"
	"strconv"
)

func markElements(root *Node, prefix string) (elements []string) {
	if root.Char != '\x00' {
		elements = append(elements, fmt.Sprintf("%s=%s", string(root.Char), prefix))
	}

	if root.Left != nil {
		elements = append(elements, markElements(root.Left, prefix+"0")...)
	}
	if root.Right != nil {
		elements = append(elements, markElements(root.Right, prefix+"1")...)
	}

	return elements
}

func main() {
	const str = "beep boop beer!"
	//tree := Tree()
	//for _, entry := range tree.AsList() {
	//	fmt.Println(entry)
	//}

	compressed, bits := Compress(str)
	for _, u32 := range compressed {
		fmt.Print(pad(strconv.FormatUint(uint64(u32), 2), "0", 32) + " ")
	}
	fmt.Println()
	fmt.Println("total bits:", bits)
}
