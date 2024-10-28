func calPoints(operations []string) int {
    var ans[]int

    for _, operation := range operations {
        if num, err := strconv.Atoi(operation); err == nil {
            ans = append(ans, num)
        } else if operation == \D\ && len(ans) > 0 {
            currentScore := ans[len(ans) - 1] * 2
            ans = append(ans, currentScore)
        } else if operation == \C\ && len(ans) > 0 {
            ans = ans[:len(ans) - 1]
        } else if operation == \+\ && len(ans) >= 2 {
            currentScore := ans[len(ans) - 1] + ans[len(ans) - 2]
            ans = append(ans, currentScore)
        }
    }

    total := 0
    for _, score := range ans {
        total += score
    }
    return total
}