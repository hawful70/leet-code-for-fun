func isValid(s string) bool {
    stack := []rune{}
    closeToOpen := map[rune]rune{')': '(', '}': '{', ']': '['}
    
    for _, c := range s {
        if open, found := closeToOpen[c]; found {
            if len(stack) > 0 && stack[len(stack) - 1] == open {
                stack = stack[:len(stack) - 1]
            } else {
                return false
            }
        } else {
            stack = append(stack, c)
        }
    }

    return len(stack) == 0
}