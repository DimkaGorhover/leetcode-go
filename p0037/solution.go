package p0037

const zeroB = byte(0)

func solveSudoku(board [][]byte) {
	newSudokuSolver(board).solve()
}

func newSudokuSolver(board [][]byte) *sudokuSolver {
	s := &sudokuSolver{
		board:     board,
		rowsSets:  &lineSet{},
		colsSets:  &lineSet{},
		squareSet: &squaresSets{},
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if num, isNum := getNum(board[row][col]); isNum {
				s.set(row, col, num)
			}
		}
	}

	return s
}

func getNum(c byte) (byte, bool) {
	return c - '1', '1' <= c && c <= '9'
}

type sudokuSolver struct {
	board     sudokuBoard
	rowsSets  *lineSet
	colsSets  *lineSet
	squareSet *squaresSets
}

func (s *sudokuSolver) solve() bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			if s.board.isSet(row, col) {
				continue
			}

			for value := zeroB; value < 9; value++ {
				if s.isFree(row, col, value) {

					s.set(row, col, value)

					if s.solve() {
						return true
					}

					s.reset(row, col, value)
				}
			}

			return false
		}
	}

	return true
}

func (s *sudokuSolver) isFree(row, col int, value byte) bool {
	return s.board.isFree(row, col) &&
		!s.rowsSets.isFree(row, value) &&
		!s.colsSets.isFree(col, value) &&
		!s.squareSet.isFree(row, col, value)
}

func (s *sudokuSolver) set(row, col int, value byte) {
	s.board.set(row, col, value)
	s.rowsSets.set(row, value)
	s.colsSets.set(col, value)
	s.squareSet.set(row, col, value)
}

func (s *sudokuSolver) reset(row, col int, value byte) {
	s.board.reset(row, col)
	s.rowsSets.reset(row, value)
	s.colsSets.reset(col, value)
	s.squareSet.reset(row, col, value)
}

type sudokuBoard [][]byte

func (b sudokuBoard) isSet(row, col int) bool {
	return !b.isFree(row, col)
}

func (b sudokuBoard) set(row, col int, value byte) {
	b[row][col] = value + '1'
}

func (b sudokuBoard) reset(row, col int) {
	b[row][col] = '.'
}

func (b sudokuBoard) isFree(row, col int) bool {
	return b[row][col] == '.'
}

type squaresSets [9][9]bool

func (s *squaresSets) isFree(row, col int, num byte) bool {
	index := (row / 3) + (col/3)*3
	return s[index][num]
}

func (s *squaresSets) set(row, col int, num byte) {
	index := (row / 3) + (col/3)*3
	s[index][num] = true
}

func (s *squaresSets) reset(row, col int, num byte) {
	index := (row / 3) + (col/3)*3
	s[index][num] = false
}

type lineSet [9][9]bool

func (s *lineSet) isFree(index int, num byte) bool {
	return s[index][num]
}

func (s *lineSet) set(index int, num byte) {
	s[index][num] = true
}

func (s *lineSet) reset(index int, num byte) {
	s[index][num] = false
}
