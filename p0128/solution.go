package p0128

func longestConsecutive(nums []int) int {
	t := map[int]bool{}
	for _, num := range nums {
		t[num] = true
	}
	res := 0
	for _, num := range nums {
		if t[num-1] {
			continue
		}
		l := 0
		for ; t[num]; num++ {
			l++
		}
		res = max(res, l)
	}
	return res
}
