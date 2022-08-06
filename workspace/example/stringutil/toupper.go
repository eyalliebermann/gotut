package stringutil

import "unicode"

func ToUpper(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}

	return string(r)
}
