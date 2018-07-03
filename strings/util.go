package strings

import (
	"github.com/satori/go.uuid"
	"regexp"
	"strings"
)

// Reverse return a reserved string
func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	return string(reverse(r))
}

// MOBILE_REGEX  define mobile regex
const MOBILE_REGEX = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9])|(17[0-9]))\\d{8}$"

// IsValidMobile returns a string is a phone No. or not
func IsValidMobile(mobileNo string) bool {
	reg := regexp.MustCompile(MOBILE_REGEX)
	return reg.MatchString(mobileNo)
}

// UUIDType define uuid display type ,int
type UUIDType int

// define the uuid types
const (
	UUID_TYPE_HASH_LIKE UUIDType = iota
	UUID_TYPE_CANONICAL
)

// NewUUID is a func that create a new uuid and returns
func NewUUID(t UUIDType) (string, error) {
	id, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	switch t {
	case UUID_TYPE_HASH_LIKE:
		return strings.Replace(id.String(), "-", "", -1), nil
	case UUID_TYPE_CANONICAL:
		return id.String(), nil
	}
	return id.String(), nil
}

func reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
