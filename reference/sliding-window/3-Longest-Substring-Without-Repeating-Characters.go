func lengthOfLongestSubstring(s string) int {
    repeatCharacter := make(map[rune]bool)
    left := 0
    ans := 0
    for right, c := range s {
        for repeatCharacter[c] {
            delete(repeatCharacter, rune(s[left]))
            left++
        }
        repeatCharacter[c] = true
        ans = max(ans, right - left + 1)
    }
    return ans
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}