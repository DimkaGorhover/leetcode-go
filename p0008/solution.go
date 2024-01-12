package p0008

import (
	"math"
)

const (
	prevMaxInt = math.MaxInt32 / 10

	preMaxIntDiv = math.MaxInt32 % 10

	space = uint8(' ')
	minus = uint8('-')
	plus  = uint8('+')
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	i := 0

	for i < len(s) && s[i] == space {
		i++
	}

	sign := 1
	if i < len(s) {
		switch s[i] {
		case minus:
			sign = -1
			i++
		case plus:
			i++
		}
	}

	result := 0
	for ; i < len(s); i++ {
		num := int(s[i]) - '0'
		if 0 > num || num > 9 {
			break
		}
		if result > prevMaxInt || prevMaxInt == result && preMaxIntDiv < num {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + num
	}

	return result * sign
}
