package lc0003

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}
	if l == 2 {
		if s[0] == s[1] {
			return 1
		}
		return 2
	}
	index := [1 << 7]int{}
	i := 0
	m := 1
	for j := 0; j < l; j++ {
		c := s[j]
		i = max(index[c], i)
		m = max(m, j-i+1)
		index[c] = j + 1
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
