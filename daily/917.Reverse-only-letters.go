package daily

import "unicode"

func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	left := 0
	right := len(s) - 1

	for left <= right {
		if !unicode.IsLetter(runes[left]) {
			left++
		} else if !unicode.IsLetter(runes[right]) {
			right--
		} else {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}

	}

	return string(runes)
}
