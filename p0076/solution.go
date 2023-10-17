package p0076

import (
	"math"
	"strconv"
	"strings"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	set := newCharsSet(t)

	sLen := len(s)
	start := 0
	end := 0
	i := 0
	j := 0
	minLength := math.MaxInt

	for j < sLen {

		for set.hasMore() && j < sLen {
			set.acquire(s[j])
			j++
		}

		if set.hasMore() && j == sLen {
			break
		}

		prevI := i
		for !set.hasMore() {
			set.release(s[i])
			i++
			prevI = i - 1
		}

		length := j - prevI
		if length < minLength {
			minLength = length
			start = prevI
			end = j
		}

	}

	return s[start:end]
}

type charsSet struct {
	set      [1 << 7]bool
	window   [1 << 7]int
	count    int
	acquires int
}

func newCharsSet(t string) *charsSet {
	set := [1 << 7]bool{}
	window := [1 << 7]int{}
	count := 0
	for i := 0; i < len(t); i++ {
		set[t[i]] = true
		window[t[i]]++
		count++
	}
	return &charsSet{set, window, count, 0}
}

func (s *charsSet) hasMore() bool {
	return s.acquires < s.count
}

func (s *charsSet) acquire(c uint8) {
	if s.set[c] {
		count := s.window[c]
		s.window[c]--
		if count > 0 {
			s.acquires++
		}
	}
}

func (s *charsSet) release(c uint8) {
	if s.set[c] {
		s.window[c]++
		count := s.window[c]
		if count > 0 {
			s.acquires--
		}
	}
}

func (s *charsSet) String() string {
	sb := strings.Builder{}
	first := true
	for i := 0; i < len(s.set); i++ {
		if s.set[i] {
			if !first {
				sb.WriteString(", ")
			}
			first = false
			sb.WriteString(strconv.Itoa(i))
			sb.WriteString("=")
			sb.WriteString(strconv.Itoa(s.window[i]))
		}
	}
	return sb.String()
}
