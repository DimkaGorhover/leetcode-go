package lc0287

var s solution = &hashSolution{}

func findDuplicate(nums []int) int {
	return s.findDuplicate(nums)
}

type solution interface {
	findDuplicate(nums []int) int
}

type hashSolution struct{}

func (s *hashSolution) findDuplicate(nums []int) int {
	set := [100_001]bool{}
	for _, num := range nums {
		if set[num] {
			return num
		}
		set[num] = true
	}
	return 0
}

type negativeArraySolution struct{}

func (s *negativeArraySolution) findDuplicate(nums []int) int {
	for _, num := range nums {
		i := absInt(num)
		if nums[i] < 0 {
			return i
		}
		nums[i] = -nums[i]
	}
	return 0
}

func absInt(n int) int {
	return (n ^ (n >> 31)) + int(uint32(n)>>31)
}
