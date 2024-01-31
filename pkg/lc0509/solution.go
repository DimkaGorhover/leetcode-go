package lc0509

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	dp := [2]int{1, 2}
	for i := 4; i <= n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}
	return dp[n%2]
}
