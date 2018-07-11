package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSundaySearch(t *testing.T) {
	s, p := "abcdabd", "bd"
	i := SundaySearch(s, p)
	assert.Equal(t, 5, i, "")
}
