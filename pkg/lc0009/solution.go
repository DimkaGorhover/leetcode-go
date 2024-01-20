package lc0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	r := 0
	tmp := x
	for tmp > 0 {
		r = (r * 10) + (tmp % 10)
		tmp = tmp / 10
	}
	return x == r
}
