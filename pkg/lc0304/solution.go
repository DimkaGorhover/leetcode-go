package lc0304

import (
	"fmt"
	"strings"
)

type NumMatrix [][]int

func Constructor(matrix [][]int) NumMatrix {

	rows := len(matrix)
	cols := len(matrix[0])

	m := make(NumMatrix, rows+1)
	m[0] = make([]int, cols+1)
	m[1] = make([]int, cols+1)

	for col := 0; col < cols; col++ {
		m[1][col+1] = m[1][col] + matrix[0][col]
	}

	for row := 1; row < rows; row++ {
		m[row+1] = make([]int, cols+1)
		for col := 0; col < cols; col++ {
			m[row+1][col+1] = m[row][col+1] + m[row+1][col] + matrix[row][col] - m[row][col]
		}
	}

	return m
}

func (m NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return m[row2+1][col2+1] - m[row1][col2+1] - m[row2+1][col1] + m[row1][col1]
}

func (m NumMatrix) String() string {
	var sb strings.Builder
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[row]); col++ {
			sb.WriteString(fmt.Sprintf(" %2v", m[row][col]))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
