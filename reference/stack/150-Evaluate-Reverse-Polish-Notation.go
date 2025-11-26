func evalRPN(tokens []string) int {
    if len(tokens) <= 1 {
        value, _ := strconv.Atoi(tokens[0])
        return value
    }

    var stacks []int
    for _, token := range tokens {
        switch token {
            case \+\:
                firstNumber := stacks[len(stacks) - 1]
                secondNumber := stacks[len(stacks) - 2]
                stacks = stacks[:len(stacks) - 2]
                stacks = append(stacks, secondNumber + firstNumber)
            case \-\:
                firstNumber := stacks[len(stacks) - 1]
                secondNumber := stacks[len(stacks) - 2]
                stacks = stacks[:len(stacks) - 2]
                stacks = append(stacks, secondNumber - firstNumber)
            case \*\:
                firstNumber := stacks[len(stacks) - 1]
                secondNumber := stacks[len(stacks) - 2]
                stacks = stacks[:len(stacks) - 2]
                stacks = append(stacks, secondNumber * firstNumber)
            case \/\:
                firstNumber := stacks[len(stacks) - 1]
                secondNumber := stacks[len(stacks) - 2]
                stacks = stacks[:len(stacks) - 2]
                stacks = append(stacks, secondNumber / firstNumber)
            default:
                number, _ := strconv.Atoi(token)
                stacks = append(stacks, number)
        }
    }

    return stacks[len(stacks) - 1]
}