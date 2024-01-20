package lc0013

func romanToInt(s string) int {

	length := len(s)
	if length == 0 {
		return 0
	}

	values := [128]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if length == 1 {
		return values[s[0]]
	}

	pairs := [128]byte{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}

	i := 0
	sum := 0
	for i < length-1 {
		c1 := s[i]
		c2 := s[i+1]
		if pairs[c2] == c1 {
			sum += values[c2] - values[c1]
			i += 2
		} else {
			sum += values[c1]
			i++
		}
	}

	if i < length {
		sum += values[s[i]]
	}

	return sum
}
