package p1122

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	table := [1001]int{}
	for i := range arr2 {
		table[arr2[i]] = i + 1
	}
	sort.Slice(arr1, func(i, j int) bool {
		if iPos := table[arr1[i]]; iPos > 0 {
			if jPos := table[arr1[j]]; jPos > 0 {
				return iPos < jPos
			}
			return true
		}

		if jPos := table[arr1[j]]; jPos > 0 {
			return false
		}

		return arr1[i] < arr1[j]
	})
	return arr1
}
