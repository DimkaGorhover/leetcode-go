package p2283

func digitCount(num string) bool {
	ht := [10]uint8{}
	n := uint8(len(num))
	var i uint8
	for i = 0; i < n; i++ {
		digit := num[i] - '0'
		ht[digit]++
	}

	for i = 0; i < n; i++ {
		digit := num[i] - '0'
		count := ht[i]
		if digit != count {
			return false
		}
	}

	return true
}
