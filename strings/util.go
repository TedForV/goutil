package strings

func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	return string(reverse(r))
}

func reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
