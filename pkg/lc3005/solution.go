package lc3005

func maxFrequencyElements(nums []int) int {
	a := [101]int{}
	b := [101]int{}
	m := 0
	for _, num := range nums {
		a[num]++
		m = max(m, a[num])
		b[a[num]]++
	}
	return m * b[m]
}
