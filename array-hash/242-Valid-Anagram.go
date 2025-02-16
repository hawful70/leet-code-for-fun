func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	character_map := make(map[rune]int)

	for _, char := range s {
		character_map[char]++
	}

	for _, char := range t {
		character_map[char]--
	}

	for _, value := range character_map {
		if value > 0 {
			return false
		}
	}

	return true
}