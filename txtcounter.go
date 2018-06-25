package main

import (
	"unicode"
	"unicode/utf8"

	"github.com/lxn/walk"
)

func CountAll(s string) int {

	walk.MsgBox(
		nil,
		"CountAll",
		"CountAll",
		walk.MsgBoxOK|walk.MsgBoxIconError)

	return len([]rune(s))
}

func CountRemoveBlank(s string) int {
	var cntBlank int
	textLength := CountAll(s)

	for i := 0; i < textLength; i++ {
		if string([]rune(s)[i]) == " " || string([]rune(s)[i]) == "\n" || string([]rune(s)[i]) == "\t" {

			cntBlank++
		}
	}
	excludeBlank := textLength - cntBlank
	return excludeBlank
}

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
