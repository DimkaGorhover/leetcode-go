package lc0345

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	b := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isVowel(b[i]) {
			i++
		}
		for i < j && !isVowel(b[j]) {
			j--
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
