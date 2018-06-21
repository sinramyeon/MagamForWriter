package main

import (
	"unicode"
	"unicode/utf8"
)

func CountAll(s string) int{
	return len([]rune(s))
}

func CountRemoveBlank(s string) int{
	var cnt_blank int
	stored_text_length := CountAll(s)

	for(int i=0; i<stored_text_length; i++){
		if(s.charAt(i) == ' ' || s.charAt(i) == '\n' || s.charAt(i) == '\t')
			cnt_blank++;
	}
	exclude_blank_words := total_count_words -  cnt_blank;
	return exclude_blank_words
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
