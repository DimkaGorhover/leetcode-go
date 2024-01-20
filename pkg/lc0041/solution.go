package lc0041

func firstMissingPositive(nums []int) int {
	ht := map[int]bool{}
	minNum := 1
	maxNum := 1
	for _, num := range nums {
		if num > 0 {
			ht[num] = true
			minNum = min(minNum, num)
			maxNum = max(maxNum, num)
		}
	}
	for num := 1; num <= maxNum; num++ {
		if _, ok := ht[num]; !ok {
			return num
		}
	}
	return maxNum + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
