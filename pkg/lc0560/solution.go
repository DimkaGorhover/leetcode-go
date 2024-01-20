package lc0560

// TODO
func subarraySum(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n)
	prev := 0
	for i := 0; i < n; i++ {
		prev += nums[i]
		prefixSum[i] = prev
	}

	return 0
}
