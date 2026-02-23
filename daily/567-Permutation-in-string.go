package daily

import "reflect"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	targetWindow := make(map[rune]int)
	for _, char := range s1 {
		targetWindow[char]++
	}
	k := len(s1)
	currentWindow := make(map[rune]int)
	for i := 0; i < k; i++ {
		currentWindow[rune(s2[i])]++
	}

	for i := k; i < len(s2); i++ {
		if reflect.DeepEqual(targetWindow, currentWindow) {
			return true
		}
		currentWindow[rune(s2[i])]++
		currentWindow[rune(s2[i-k])]--
		if currentWindow[rune(s2[i-k])] == 0 {
			delete(currentWindow, rune(s2[i-k]))
		}
	}
	return reflect.DeepEqual(targetWindow, currentWindow)
}
