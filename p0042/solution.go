package p0042

func trap(height []int) int {
	leftIndex := 0
	rightIndex := len(height) - 1
	level := 0
	water := 0
	lower := 0
	for leftIndex < rightIndex {
		left := height[leftIndex]
		right := height[rightIndex]
		if left < right {
			lower = left
			leftIndex++
		} else {
			lower = right
			rightIndex--
		}
		level = maxInt(level, lower)
		water += level - lower
	}
	return water
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
