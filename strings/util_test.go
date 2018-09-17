package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	s := "shit world和大家"
	result := Reverse(s)
	assert.Equal(t, "家大和dlrow tihs", result, "they should be reversed")
}

func TestIsValidMobile(t *testing.T) {
	s := "18503060972"
	assert.Equal(t, true, IsValidMobile(s), "is valid")
	s = "185030609721"
	assert.Equal(t, false, IsValidMobile(s), "is not valid")
}

func TestNewUUID(t *testing.T) {
	s, err := NewUUID(UUIDTypeCanonical)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s)
	s, err = NewUUID(UUIDTypeHashLike)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s)
}
