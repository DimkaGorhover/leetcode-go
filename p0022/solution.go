package p0022

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	if n == 2 {
		return []string{
			"(())",
			"()()",
		}
	}

	chars := make(chars, n<<1)
	chars[0] = '('

	return doGenerate([]string{}, n-1, n, chars, 1)
}

type chars []byte

func (c chars) String() string {
	return string(c)
}

func doGenerate(list []string, open, close int, chars chars, charsIndex int) []string {
	if open > 0 || close > 0 {
		if open > 0 {
			chars[charsIndex] = '('
			list = doGenerate(list, open-1, close, chars, charsIndex+1)
		}
		if close > 0 && open < close {
			chars[charsIndex] = ')'
			list = doGenerate(list, open, close-1, chars, charsIndex+1)
		}
	} else {
		list = append(list, chars.String())
	}
	return list
}
