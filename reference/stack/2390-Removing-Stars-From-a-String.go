func removeStars(s string) string {
    var stacks []rune
    for _, c := range s {
        if len(stacks) > 0 && c == '*' && stacks[len(stacks) - 1] != '*' {
            stacks = stacks[:len(stacks) - 1]
        } else {
            stacks = append(stacks, c)
        }
    }

    return string(stacks)    
}