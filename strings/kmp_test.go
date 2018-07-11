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

func TestGetImprovedNext(t *testing.T) {
	s := "abcdabd" //[-1,0,0,0,0,1,2]
	r := []rune(s)
	result := GetImprovedNext(r)
	assert.Equal(t, result[0], -1, "")
	assert.Equal(t, result[1], 0, "")
	assert.Equal(t, result[2], 0, "")
	assert.Equal(t, result[3], 0, "")
	assert.Equal(t, result[4], -1, "")
	assert.Equal(t, result[5], 0, "")
	assert.Equal(t, result[6], 2, "")
}

func TestKMPImprovedSearch(t *testing.T) {
	s, p := "abcdabd", "bd"
	i := KMPImprovedSearch(s, p)
	assert.Equal(t, 5, i, "")
}

func BenchmarkKMPSearch(b *testing.B) {
	s, p := "abcdawefasdfcasefasdvzxcvawesdfgsfdbd", "bd"
	for i := 0; i < b.N; i++ {
		KMPSearch(s, p)
	}
}

func BenchmarkKMPImprovedSearch(b *testing.B) {
	s, p := "abcdawefasdfcasefasdvzxcvawesdfgsfdbd", "bd"
	for i := 0; i < b.N; i++ {
		KMPSearch(s, p)
	}
}
