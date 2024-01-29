package lc0258

func addDigits(num int) int {
	for num > 9 {
		num = sum(num)
	}
	return num
}

func sum(num int) int {
	s := 0
	for num > 0 {
		s = s + (num % 10)
		num /= 10
	}
	return s
}
