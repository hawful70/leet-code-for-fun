package daily

func maxFreqSum(s string) int {
	maxVowel, maxConsonant := 0, 0
	map_Vowel, map_consonant := make(map[rune]int), make(map[rune]int)

	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			map_Vowel[rune(c)]++
			maxVowel = max(maxVowel, map_Vowel[rune(c)])
		} else {
			map_consonant[rune(c)]++
			maxConsonant = max(maxConsonant, map_consonant[rune(c)])
		}
	}

	return maxVowel + maxConsonant
}
