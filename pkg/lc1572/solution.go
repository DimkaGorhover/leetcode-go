package lc1572

func diagonalSum(mat [][]int) int {
	n := len(mat)
	sum := 0
	for i, row := range mat {
		j := n - i - 1
		sum += row[j]
		if i != j {
			sum += row[n-1-j]
		}
	}
	return sum
}
