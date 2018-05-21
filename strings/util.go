package strings

import "regexp"

func Reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	return string(reverse(r))
}

const MOBILE_REGEX = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"

func IsValidMobile(mobileNo string) bool {
	reg := regexp.MustCompile(MOBILE_REGEX)
	return reg.MatchString(mobileNo)
}

func reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
