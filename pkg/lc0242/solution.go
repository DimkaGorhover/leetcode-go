package lc0242

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	table := ['z' - 'a' + 1]int{}
	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
		table[t[i]-'a']--

	}
	for _, v := range table {
		if v < 0 {
			return false
		}
	}

	return true
}
