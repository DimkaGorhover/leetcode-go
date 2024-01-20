package lc0055

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_canJump(t *testing.T) {
	t.Parallel()

	type args struct {
		nums []int
	}
	numbers, expected := readTestCase()
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test_001", args: args{nums: []int{0}}, want: true},
		{name: "test_002", args: args{nums: []int{0, 2, 3}}, want: false},
		{name: "test_003", args: args{nums: []int{3, 2, 1, 0, 4}}, want: false},
		{name: "test_004", args: args{nums: []int{2, 3, 1, 1, 4}}, want: true},
		{name: "test_005", args: args{nums: numbers}, want: expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func readTestCase() ([]int, bool) {
	file, err := os.Open("test_case_001.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	str = str[:len(str)-1]
	arr := strings.Split(str, ",")
	numbers := make([]int, len(arr), len(arr))
	for i, str := range arr {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		numbers[i] = number
	}
	str, err = reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	expected, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return numbers, expected
}
