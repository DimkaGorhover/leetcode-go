package lc0011

import "math"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	if len(height) == 2 {
		return min(height[0], height[1])
	}
	i := 0
	j := len(height) - 1
	m := math.MinInt
	for i < j {
		h := min(height[i], height[j])
		if current := (j - i) * h; current > m {
			m = current
		}
		for height[i] <= h && i < j {
			i++
		}
		for height[j] <= h && i < j {
			j--
		}
	}
	return m
}
