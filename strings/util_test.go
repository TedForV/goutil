package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	s := "shit world和大家"
	result := Reverse(s)
	assert.Equal(t, "家大和dlrow tihs", result, "they should be reversed")
}
