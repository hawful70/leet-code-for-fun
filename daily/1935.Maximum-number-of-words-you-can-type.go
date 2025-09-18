package daily

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	brokenMap := make(map[rune]bool)
	for _, c := range brokenLetters {
		brokenMap[rune(c)] = true
	}

	res := 0
	words := strings.Split(text, " ")
	for _, word := range words {
		for _, c := range word {
			if brokenMap[rune(c)] {
				res++
				break
			}
		}
	}

	return len(words) - res
}
