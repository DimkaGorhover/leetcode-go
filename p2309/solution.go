package p2309

func greatestLetter(s string) string {
	if len(s) <= 1 {
		return ``
	}
	table := make([]int, 26, 26)
	var maxIndex int32 = -1
	for _, charCode := range s {
		index := charCode - 65
		if index > 31 {
			// lower case
			index = index - 32
			table[index] = table[index] | 1
		} else {
			// upper case
			table[index] = table[index] | 2
		}
		if table[index] == 3 && index > maxIndex {
			maxIndex = index
		}
	}
	if maxIndex >= 0 {
		return string(maxIndex + 65)
	}
	return ""
}
