package p2133

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	if n == 0 {
		return true
	}

	if n == 1 {
		if matrix[0][0] == 1 {
			return true
		}
		return false
	}

	set := newSetMatrix(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := matrix[i][j]
			set.addRow(i, num)
			set.addCol(j, num)
		}
	}

	return set.isEmpty()
}

type setMatrix struct {
	rows     [][]bool
	cols     [][]bool
	count    int
	capacity int
}

func newSetMatrix(capacity int) *setMatrix {
	return &setMatrix{
		rows:     make([][]bool, capacity, capacity),
		cols:     make([][]bool, capacity, capacity),
		count:    capacity * capacity * 2,
		capacity: capacity,
	}
}

func (s *setMatrix) addRow(row, num int) {
	num -= 1
	if num < 0 || num >= s.capacity {
		return
	}
	if s.rows[row] == nil {
		s.rows[row] = make([]bool, s.capacity, s.capacity)
	}
	if s.rows[row][num] {
		return
	}
	s.rows[row][num] = true
	s.count--
}

func (s *setMatrix) addCol(col, num int) {
	num -= 1
	if num < 0 || num >= s.capacity {
		return
	}
	if s.cols[col] == nil {
		s.cols[col] = make([]bool, s.capacity, s.capacity)
	}
	if s.cols[col][num] {
		return
	}
	s.cols[col][num] = true
	s.count--
}

func (s *setMatrix) isEmpty() bool {
	return s.count == 0
}
