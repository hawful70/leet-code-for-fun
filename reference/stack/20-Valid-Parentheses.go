func isValid(s string) bool {
    closeToOpen := map[rune]rune { ')': '(', '}': '{', ']': '[' }
    var stacks []rune
    for _, c := range s {
        if open, found := closeToOpen[c]; found {
            if len(stacks) > 0 && open == stacks[len(stacks) - 1] {
                stacks = stacks[:len(stacks) - 1]
            } else {
                return false
            }
        } else {
            stacks = append(stacks, c)
        }
        
    }
    return len(stacks) == 0
}