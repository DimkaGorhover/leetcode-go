package p1137

func tribonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	arr := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		arr[i%3] = arr[0] + arr[1] + arr[2]
	}
	return arr[n%3]
}
