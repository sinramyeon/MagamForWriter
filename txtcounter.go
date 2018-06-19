package main

import (
	"unicode"
	"unicode/utf8"
)

// https://engineering.linecorp.com/ko/blog/detail/52
func CountChar(s string) int {
	if len(s) == 0 {
		return 0
	}
	gr := 1
	_, s1 := utf8.DecodeRuneInString(s)
	for _, r := range s[s1:] {
		if !unicode.Is(unicode.Mn, r) {
			gr++
		}
	}
	return gr
}
