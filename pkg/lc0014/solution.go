package lc0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	head := strs[0]
	if len(strs) == 1 {
		return head
	}

	for i := 1; i < len(strs); i++ {
		head = getPrefix(head, strs[i])
		if len(head) == 0 {
			return ""
		}
	}

	return head
}

func getPrefix(s1, s2 string) string {
	length := min(len(s1), len(s2))
	var prefix []byte
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			break
		}
		prefix = append(prefix, s1[i])
	}

	return string(prefix)
}
