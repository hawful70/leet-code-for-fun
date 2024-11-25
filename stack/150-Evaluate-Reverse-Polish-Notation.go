func evalRPN(tokens []string) int {
    var stacks []int
    if len(tokens) <= 1 {
        number, _ := strconv.Atoi(tokens[0])
        return number
    }

    for _, c := range tokens {
        switch c {
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
            number, _ := strconv.Atoi(c)
            stacks = append(stacks, number)
        } 
    }
    return stacks[len(stacks) - 1]
}