package lc2108

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(text string) bool {
	i := 0
	j := len(text) - 1
	for i < j {
		if text[i] != text[j] {
			return false
		}
		i++
		j--
	}
	return true
}
