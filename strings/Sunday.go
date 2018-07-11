package strings

func SundaySearch(content string, pattern string) int {
	cr, pr := []rune(content), []rune(pattern)
	ci, pi, cl, pl := 0, 0, len(cr), len(pr)
	pmap := make(map[rune]int, pl)
	for i, v := range pr {
		if index, ok := pmap[v]; ok {
			if i < index {
				pmap[v] = i
			}
		} else {
			pmap[v] = i
		}
	}

	for ci < cl && pi < pl {
		if cr[ci] == pr[pi] {
			ci, pi = ci+1, pi+1
		} else {
			if v, ok := pmap[cr[ci]]; ok {
				ci, pi = ci+pl-v+1, 0
			} else {
				ci, pi = ci+pl+1, 0
			}
		}
	}
	if pi == pl {
		return ci - pi
	} else {
		return -1
	}
}
