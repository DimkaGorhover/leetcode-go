package p0004

import "math"

// LeetCode #4: Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	getPos := func(arr []int, pos int) int {
		if pos < len(arr) {
			return arr[pos]
		}
		return math.MaxInt
	}

	sumLen := len(nums1) + len(nums2)
	newLen := (sumLen >> 1) + 1
	ints := make([]int, newLen, newLen)
	i1 := 0
	i2 := 0
	for i := 0; i < newLen; i++ {
		num1 := getPos(nums1, i1)
		num2 := getPos(nums2, i2)
		if num1 < num2 {
			ints[i] = num1
			i1++
		} else {
			ints[i] = num2
			i2++
		}
	}

	if sumLen%2 == 0 {
		return 0.5 * float64(ints[len(ints)-2]+ints[len(ints)-1])
	}

	return float64(ints[len(ints)-1])
}
