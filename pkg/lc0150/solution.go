package lc0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	s := &stack{}
	for _, token := range tokens {
		if len(token) == 1 {
			oneElement(s, token[0])
		} else {
			number, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			s.push(number)
		}
	}
	return s.pop()
}

func oneElement(s *stack, c uint8) {
	switch c {
	case '+':
		v1 := s.pop()
		v2 := s.pop()
		s.push(v2 + v1)
	case '-':
		v1 := s.pop()
		v2 := s.pop()
		s.push(v2 - v1)
	case '*':
		v1 := s.pop()
		v2 := s.pop()
		s.push(v2 * v1)
	case '/':
		v1 := s.pop()
		v2 := s.pop()
		s.push(v2 / v1)
	default:
		s.push(int(c - '0'))
	}
}

type stack struct {
	arr    []int
	cursor int
}

func (s *stack) push(value int) {
	if s.cursor < len(s.arr) {
		s.arr[s.cursor] = value
	} else {
		s.arr = append(s.arr, value)
	}
	s.cursor++
}

func (s *stack) pop() int {
	if s.cursor == 0 {
		panic("stack is empty")
	}
	s.cursor--
	value := s.arr[s.cursor]
	return value
}
