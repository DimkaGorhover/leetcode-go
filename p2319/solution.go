package p2319

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i == n-1-j {
				if grid[i][j] == 0 {
					return false
				}
			} else if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
