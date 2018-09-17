package strings

// GetNext is get the next array of KMP
func GetNext(pattern []rune) []int {
	l := len(pattern)
	next := make([]int, l)
	next[0] = -1
	k, j := -1, 0
	for j < l-1 {
		if k == -1 || pattern[j] == pattern[k] {
			k, j = k+1, j+1
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

// GetImprovedNext improved method for Next
func GetImprovedNext(pattern []rune) []int {
	l := len(pattern)
	next := make([]int, l)
	next[0] = -1
	k, j := -1, 0
	for j < l-1 {
		if k == -1 || pattern[j] == pattern[k] {
			k, j = k+1, j+1
			if pattern[j] != pattern[k] {
				next[j] = k
			} else {
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
	return next
}

// KMPSearch search substring by KMP algorithm
func KMPSearch(content string, pattern string) int {
	cr, pr := []rune(content), []rune(pattern)
	next := GetNext(pr)
	i, j, sl, pl := 0, 0, len(cr), len(pr)
	for i < sl && j < pl {
		if j == -1 || cr[i] == pr[j] {
			i, j = i+1, j+1
		} else {
			j = next[j]
		}
	}
	if j == pl {
		return i - j
	}
	return -1

}

// KMPImprovedSearch improved kmp search
func KMPImprovedSearch(content string, pattern string) int {
	cr, pr := []rune(content), []rune(pattern)
	next := GetImprovedNext(pr)
	i, j, sl, pl := 0, 0, len(cr), len(pr)
	for i < sl && j < pl {
		if j == -1 || cr[i] == pr[j] {
			i, j = i+1, j+1
		} else {
			j = next[j]
		}
	}
	if j == pl {
		return i - j
	}
	return -1

}
