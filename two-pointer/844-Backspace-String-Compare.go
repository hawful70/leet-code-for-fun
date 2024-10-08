func backspaceCompare(s string, t string) bool {
    stackStringCompare := func(str string) []rune {
        stack := []rune{}
        for _, char := range str {
            if char != '#' {
                stack = append(stack, char)
            } else if len(stack) > 0 {
                stack = stack[:len(stack) - 1]
            }
        }
        return stack
    }

    return string(stackStringCompare(s)) == string(stackStringCompare(t))
}