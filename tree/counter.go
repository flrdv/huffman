package tree

func countLetters(str string) (nodes []weightedNode) {
	for i := 0; i < len(str); i++ {
		index := nodeIndex(str[i], nodes)
		if index == -1 {
			nodes = append(nodes, newWeightedNode(str[i], 1, nil, nil))
		} else {
			nodes[index].Frequency++
		}
	}

	return nodes
}

func nodeIndex(char byte, nodes []weightedNode) int {
	for i, node := range nodes {
		if node.Char == char {
			return i
		}
	}

	return -1
}
