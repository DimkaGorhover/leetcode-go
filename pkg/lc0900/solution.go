package lc0900

type RLEIterator struct {
	encoding []int
	cursor   int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding, 0}
}

func (i *RLEIterator) Next(n int) int {
	for i.cursor < len(i.encoding) {
		if n > i.encoding[i.cursor] {
			n -= i.encoding[i.cursor]
			i.cursor += 2
			continue
		}
		if n == i.encoding[i.cursor] {
			i.cursor += 2
			return i.encoding[i.cursor-1]
		}
		i.encoding[i.cursor] -= n
		return i.encoding[i.cursor+1]
	}
	return -1
}
