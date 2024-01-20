package lc0205

func isIsomorphic(s string, t string) bool {
	f := [128]uint8{}
	b := [128]uint8{}
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]
		f[c1] = c2
		b[c2] = c1
	}
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]
		if c1 != b[c2] || c2 != f[c1] {
			return false
		}
	}
	return true
}
