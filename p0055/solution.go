package p0055

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	m := nums[0]
	for i := 1; i < len(nums)-1 && m > 0; i++ {
		m = maxInt(m-1, nums[i])
	}
	return m > 0
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
