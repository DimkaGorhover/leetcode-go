package p0021

//https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	table := map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]uint8, len(s))

	head := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		mapped := table[c]

		if mapped == 0 {
			stack[head] = c
			head++
			continue
		}

		top := uint8(0)
		if head > 0 {
			head--
			top = stack[head]
		}
		if top != mapped {
			return false
		}
	}
	return head == 0
}
