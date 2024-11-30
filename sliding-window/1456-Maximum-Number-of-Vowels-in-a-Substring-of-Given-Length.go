func maxVowels(s string, k int) int {
    vowelLetters := map[string]int{
        "a": 1,
        "e": 1,
        "i": 1,
        "o": 1,
        "u": 1,
    }

    total := 0
    maxNumber, left := 0, 0
    for right, _ := range s {
        if _, found := vowelLetters[string(s[right])]; found {
            total += 1
        }

        if right - left + 1 > k {
            if _, found := vowelLetters[string(s[left])]; found {
                total -= 1
            }
            left++
        }

        maxNumber = max(maxNumber, total)
    }
    return maxNumber
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}