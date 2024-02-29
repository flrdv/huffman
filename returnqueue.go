package huffman

type returnQueue struct {
	values []Symbol
	index  int
}

func newReturnQueue() *returnQueue {
	return &returnQueue{
		values: make([]Symbol, 1),
		index:  1,
	}
}

func (r *returnQueue) Push(symbols ...Symbol) {
	r.values = append(r.values, symbols...)
}

func (r *returnQueue) Pop() Symbol {
	s := r.values[r.index]
	r.index++

	return s
}

func (r *returnQueue) Return(symbol Symbol) {
	r.index--
	r.values[r.index] = symbol
}

func (r *returnQueue) IsEmpty() bool {
	return len(r.values) <= r.index
}
