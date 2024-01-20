package lc2404

func mostFrequentEven(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	ht := map[int]int{}
	for _, num := range nums {
		if num%2 == 0 {
			ht[num]++
		}
	}
	if len(ht) == 0 {
		return -1
	}
	maxFreq := 0
	minNumber := 3000
	for number, freq := range ht {
		if freq > maxFreq || (freq == maxFreq && minNumber > number) {
			maxFreq = freq
			minNumber = number
		}
	}
	return minNumber
}
