package lc0070

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	dp := [2]int{1, 2}
	for i := 2; i < n; i++ {
		dp[i%2] = dp[0] + dp[1]
	}
	return dp[(n-1)%2]
}
