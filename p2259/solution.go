package p2259

func removeDigit(number string, digit byte) string {
	n := len(number)
	pos := 0
	for i := 0; i < n; i++ {
		if number[i] == digit {
			pos = i
			if i+1 < n && number[i+1] > digit {
				break
			}
		}
	}

	chars := make([]byte, n-1)
	j := 0
	for i := 0; i < n; i++ {
		if i != pos {
			chars[j] = number[i]
			j++
		}
	}

	return string(chars)
}
