package lc0046

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	t.Parallel()
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: `test_001`,
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: `test_002`,
			args: args{
				nums: []int{},
			},
			want: [][]int{},
		},
		{
			name: `test_003`,
			args: args{
				nums: []int{1},
			},
			want: [][]int{{1}},
		},
		{
			name: `test_004`,
			args: args{
				nums: []int{1, 2},
			},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := permute(tt.args.nums)
			assert.Equal(t, len(tt.want), len(actual))
			for _, permutation := range tt.want {
				assert.Contains(t, actual, permutation)
			}
		})
	}
}

func Test_fact(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `test_001`,
			args: args{n: 2},
			want: 2,
		},
		{
			name: `test_002`,
			args: args{n: 3},
			want: 6,
		},
		{
			name: `test_003`,
			args: args{n: 4},
			want: 24,
		},
		{
			name: `test_004`,
			args: args{n: 5},
			want: 120,
		},
		{
			name: `test_005`,
			args: args{n: 6},
			want: 720,
		},
		{
			name: `test_006`,
			args: args{n: 7},
			want: 5040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fact(tt.args.n))
		})
	}
}

func Test_newLinkedList(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		l := NewLinkedList()
		_, exists := l.PollFirst()
		assert.False(t, exists)
		assert.Equal(t, "[]", fmt.Sprint(l))

		l.AddLast(1)
		l.AddLast(2)
		l.AddLast(3)
		assert.Equal(t, "[1,2,3]", fmt.Sprint(l))

		l.AddFirst(3)
		assert.Equal(t, "[3,1,2,3]", fmt.Sprint(l))

		value, exists := l.PollFirst()
		assert.True(t, exists)
		assert.Equal(t, 3, value)
		assert.Equal(t, "[1,2,3]", fmt.Sprint(l))
	})
}

func Test_permutations_String(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		p := permutations{
			{1, 2, 3},
			{1, 3, 2},
			{1, 2, 3},
		}
		assert.Equal(t, `[[1,2,3],[1,3,2],[1,2,3]]`, fmt.Sprint(p))
	})

}
