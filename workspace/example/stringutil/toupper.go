package stringutil

import "unicode"

func ToUpper(s string) string {
	if s == "" {
		return ""
	}
	// arr := []rune(s)
	// arr[0] = unicode.ToUpper(arr[0])
	// return string(arr[0])
	r := []rune(s)
	return string(unicode.ToUpper(r[0]))
}
