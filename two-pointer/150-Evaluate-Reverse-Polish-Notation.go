func evalRPN(tokens []string) int {
    stacks := make([]int, 0)

    for _, token := range tokens {
        switch token {
        case \+\:
            a := stacks[len(stacks) - 1]
            b := stacks[len(stacks) - 2]
            stacks = stacks[:len(stacks) - 2]
            stacks = append(stacks, a + b)
        case \-\:
            a := stacks[len(stacks) - 1]
            b := stacks[len(stacks) - 2]
            stacks = stacks[:len(stacks) - 2]
            stacks = append(stacks, b - a)
        case \*\:
            a := stacks[len(stacks) - 1]
            b := stacks[len(stacks) - 2]
            stacks = stacks[:len(stacks) - 2]
            stacks = append(stacks, a * b)
        case \/\:
            a := stacks[len(stacks) - 1]
            b := stacks[len(stacks) - 2]
            stacks = stacks[:len(stacks) - 2]
            stacks = append(stacks, b / a)
        default:
            num, _ := strconv.Atoi(token)
            stacks = append(stacks, num)
        }
    }
    return stacks[len(stacks) - 1]
}