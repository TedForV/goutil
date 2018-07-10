package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNext(t *testing.T) {
	s := "abcdabd" //[-1,0,0,0,0,1,2]
	r := []rune(s)
	result := GetNext(r)
	assert.Equal(t, result[0], -1, "")
	assert.Equal(t, result[1], 0, "")
	assert.Equal(t, result[2], 0, "")
	assert.Equal(t, result[3], 0, "")
	assert.Equal(t, result[4], 0, "")
	assert.Equal(t, result[5], 1, "")
	assert.Equal(t, result[6], 2, "")
}

func TestKMPSearch(t *testing.T) {
	s, p := "abcdabd", "bd"
	i := KMPSearch(s, p)
	assert.Equal(t, 5, i, "")
}
