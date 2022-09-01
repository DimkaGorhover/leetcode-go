package p0005

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	l := 0
	start := 0
	for i := 0; i < len(s)-1; i++ {
		prev := l
		l = max(l, max(expandPalindrome(s, i, i), expandPalindrome(s, i, i+1)))
		if l != prev {
			start = i - ((l - 1) >> 1)
		}
	}

	return s[start:(start + l)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func expandPalindrome(s string, left, right int) int {
	for left >= 0 && right < len(s) && left <= right && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
