func firstPalindrome(words []string) string {
    for _, word := range words {
		if util_Palindrome(word) {
            return word
        }
	}
    return ""
}

func util_Palindrome(word string) bool {
    low, high := 0, len(word) - 1
    for low < high {
        if word[low] != word[high] {
            return false
        }
        low += 1
        high -= 1
    }
    return true
}