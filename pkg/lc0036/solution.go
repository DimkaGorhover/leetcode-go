package lc0036

func isValidSudoku(board [][]byte) bool {

	sSets := squaresSets{}

	for row := 0; row < 9; row++ {

		rowSet := set{}
		colSet := set{}

		for col := 0; col < 9; col++ {

			if num, isNum := getNum(board[row][col]); isNum {
				if rowSet.getAndSet(num) || sSets.getAndSet(row, col, num) {
					return false
				}
			}

			if num, isNum := getNum(board[col][row]); isNum {
				if colSet.getAndSet(num) {
					return false
				}
			}
		}
	}

	return true
}

type squaresSets [9][9]bool

func (s *squaresSets) getAndSet(row, col int, num byte) bool {
	index := (row / 3) + (col/3)*3
	prev := s[index][num]
	s[index][num] = true
	return prev
}

type set [9]bool

func (s *set) getAndSet(num byte) bool {
	prev := s[num]
	s[num] = true
	return prev
}

func getNum(c byte) (byte, bool) {
	return c - '1', '1' <= c && c <= '9'
}
