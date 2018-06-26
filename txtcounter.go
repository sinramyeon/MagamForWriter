package main

import (
	"unicode"
	"unicode/utf8"
)

// CountAll...
// 글자수 전체 세기
func CountAll(s string) int {

	return len([]rune(s))
}

// CountRemoveBlank
// 공백 제거하고 글자수 세기
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

// CountChar ...
// 라인 블로그에서 가져온 유니코드식 글자세기라는데 한글 하나도안맞음 ㅡㅡ
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
