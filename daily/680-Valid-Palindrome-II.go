func validPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left <= right {
        if s[left] != s[right] {
            return validPalindromeUtil(s, left+1, right) || validPalindromeUtil(s, left, right-1)
        } else {
            left++
            right--
        }
    }
    return true
}

func validPalindromeUtil(s string, l, r int) bool {
    for l <= r {
        if s[l] != s[r] {
            return false
        } else {
            l++
            r--
        }
    }
    return true
}