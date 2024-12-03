func maxVowels(s string, k int) int {
    left := 0
    res := 0
    current_vowel := 0

    for right, _ := range s {
        if isVowel(rune(s[right])) {
            if isVowel(rune(s[right])) {
                current_vowel++
            }
        }

        for right - left + 1 > k {
            if isVowel(rune(s[left])) {
                current_vowel--
            }
            left++
        }

        res = max(res, current_vowel)
    }
    return res
}

func isVowel(c rune) bool {
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
        return true
    }
    return false
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}