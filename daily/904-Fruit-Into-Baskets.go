package main

func totalFruit(fruits []int) int {
	type_fruit := make(map[int]int)
	res, left := 0, 0
	for right, _ := range fruits {
		type_fruit[fruits[right]]++

		if len(type_fruit) > 2 {
			type_fruit[fruits[left]]--

			if type_fruit[fruits[left]] == 0 {
				delete(type_fruit, fruits[left])
			}
			left++
		}

		res = max(res, right-left+1)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
