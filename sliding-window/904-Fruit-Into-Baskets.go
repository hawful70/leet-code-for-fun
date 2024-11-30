func totalFruit(fruits []int) int {
    trees := make(map[int]int, 3)

    left, maxFruits := 0, 0
    
    for right, fruit := range fruits {
        if _, exist := trees[fruit]; !exist {
            trees[fruit] = 1
        } else {
            trees[fruit]++
        }

        for len(trees) > 2 {
            trees[fruits[left]]--
            if trees[fruits[left]] == 0 {
                delete(trees, fruits[left])
            }
            left++
        }

        maxFruits = max(maxFruits, right - left + 1)
    }

    return maxFruits
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}