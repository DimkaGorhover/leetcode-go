package lc0064

func minPathSum(grid [][]int) int {
	if grid == nil {
		panic("grid NPE")
	}

	rows := len(grid)
	if rows == 0 {
		panic("empty rows")
	}

	cols := len(grid[0])
	if cols == 0 {
		panic("empty cols")
	}

	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < cols; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] = min(grid[i][j]+grid[i-1][j], grid[i][j]+grid[i][j-1])
		}
	}

	return grid[rows-1][cols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
