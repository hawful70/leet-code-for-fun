func makeGood(s string) string {
    if len(s) <= 1 {
        return s
    }

    var stacks []rune
    for _, c := range s {
        if len(stacks) > 0 && c != stacks[len(stacks) - 1] && areEqualIgnoreCase(c, stacks[len(stacks) - 1]) {
            stacks = stacks[:len(stacks) - 1]
        } else {
            stacks = append(stacks, c)
        }
    }
    return string(stacks)
}

func areEqualIgnoreCase(a rune, b rune) bool {
    return unicode.ToLower(a) == unicode.ToLower(b)
}
