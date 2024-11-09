func characterReplacement(s string, k int) int {
    repeating_character := make(map[rune]int)
    ans := 0
    left := 0
    maxR := 0
    for right, c := range s {
        repeating_character[c] += 1
        maxR = max(maxR, repeating_character[c])

        if (right - left + 1) - maxR > k {
            repeating_character[rune(s[left])] -= 1
            left++
            maxR--
        }

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