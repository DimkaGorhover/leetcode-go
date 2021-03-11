package p0001

// https://leetcode.com/problems/two-sum
func twoSum(nums []int, target int) []int {
	table := make(map[int]int)

	for i, num := range nums {
		index := table[target-num]
		if index != 0 {
			return []int{index - 1, i}
		}
		table[num] = i + 1
	}

	return []int{-1, -1}
}
