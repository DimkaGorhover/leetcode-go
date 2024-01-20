package lc0037

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func cloneBoard(b [][]byte) [][]byte {
	newBoard := make([][]byte, 9)
	for row := 0; row < 9; row++ {
		newBoard[row] = make([]byte, 9)
		for col := 0; col < 9; col++ {
			newBoard[row][col] = b[row][col]
		}
	}
	return newBoard
}

func Test_solveSudoku(t *testing.T) {
	t.Parallel()
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			args: args{
				board: [][]byte{
					{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			want: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}
	for _, tt := range tests {
		board := cloneBoard(tt.args.board)
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(board)
			assert.Equal(t, tt.want, board)
		})
	}
}
