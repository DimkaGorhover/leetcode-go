package lc0071

func simplifyPath(path string) string {
	st := newStack()
	spl := newSplitter(path)

	for {
		dir, found := spl.next()
		if !found {
			break
		}
		switch dir {
		case ".", "":
			continue
		case "..":
			st.pop()
			continue
		default:
			st.push(dir)
		}
	}

	return st.String()
}

func newStack() *stack {
	return &stack{}
}

type stack struct {
	arr    []string
	cursor int
}

func (s *stack) String() string {
	if s.cursor == 0 {
		return "/"
	}
	result := ""
	for i := 0; i < s.cursor; i++ {
		result += "/" + s.arr[i]
	}
	return result
}

func (s *stack) push(value string) {
	if s.cursor < len(s.arr) {
		s.arr[s.cursor] = value
	} else {
		s.arr = append(s.arr, value)
	}
	s.cursor++
}

func (s *stack) pop() (string, bool) {
	if s.cursor == 0 {
		return "", false
	}
	s.cursor--
	value := s.arr[s.cursor]
	return value, true
}

type splitter struct {
	value  string
	cursor int
}

func newSplitter(value string) *splitter {
	return &splitter{
		value:  value,
		cursor: 1,
	}
}

func (s *splitter) next() (string, bool) {
	if s.cursor >= len(s.value) {
		return "", false
	}

	var buf []uint8
	for s.cursor < len(s.value) {
		if s.value[s.cursor] != '/' {
			buf = append(buf, s.value[s.cursor])
			s.cursor++
		} else {
			s.cursor++
			break
		}
	}

	return string(buf), s.cursor < len(s.value) || len(buf) > 0
}
