package main

import "strings"

func wordPattern(pattern string, s string) bool {
	xMap := map[byte]string{}
	yMap := map[string]byte{}
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		if yMap[word] > 0 && yMap[word] != ch || xMap[ch] != "" && xMap[ch] != word {
			return false
		}
		xMap[ch] = word
		yMap[word] = ch
	}
	return true
}
