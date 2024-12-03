func totalFruit(fruits []int) int {
    type_fruit := make(map[int]int)
    res, left := 0, 0
    for right, _ := range fruits {
        if _, found := type_fruit[fruits[right]]; found {
            type_fruit[fruits[right]]++
        } else {
            type_fruit[fruits[right]] = 1
        }

        if len(type_fruit) > 2 {
            type_fruit[fruits[left]]--

            if type_fruit[fruits[left]] == 0 {
                delete(type_fruit, fruits[left])
            }
            left++
        }

        res = max(res, right - left + 1)
    }
    return res
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}