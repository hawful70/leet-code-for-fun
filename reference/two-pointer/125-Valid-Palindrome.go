import (
\t\fmt\
\t\unicode\
    \strings\
)

func isPalindrome(s string) bool {
    left, right := 0, len(s) - 1
    for left <= right {
        if !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
            left += 1
        } else if !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
            right -= 1
        } else {
            if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
                return false
            }
            left += 1
            right -= 1
        }
    }
    return true
}