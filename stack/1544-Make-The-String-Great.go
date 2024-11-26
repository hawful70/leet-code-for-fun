func makeGood(s string) string {
    var stacks []rune
    if len(s) == 1 {
        return s
    }

    for _, c := range s {
        if len(stacks) > 0 && stacks[len(stacks) - 1] != c && areEqual(stacks[len(stacks) - 1], c) {
            stacks = stacks[:len(stacks) - 1]
        } else {
            stacks = append(stacks, c)
        }
    }

    return string(stacks)
}

func areEqual(a rune, b rune) bool {
    return unicode.ToLower(a) == unicode.ToLower(b)
}