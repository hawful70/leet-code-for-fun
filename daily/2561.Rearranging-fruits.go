package daily

import (
	"math"
	"sort"
)

func minCostRearrangingFruits(basket1 []int, basket2 []int) int64 {
	freq1 := map[int]int{}
	freq2 := map[int]int{}
	allFreq := map[int]int{}
	// step 1: Count the number of each fruit in both baskets
	for _, x := range basket1 {
		freq1[x]++
		allFreq[x]++
	}
	for _, x := range basket2 {
		freq2[x]++
		allFreq[x]++
	}

	// step 2: Check if each fruit's total count is even
	for _, v := range allFreq {
		if v%2 != 0 {
			return -1
		}
	}

	// step 3: Calculate how many fruits are “extra” in each basket
	swapFrom1 := []int{}
	swapFrom2 := []int{}
	for fruit, _ := range allFreq {
		diff := freq1[fruit] - freq2[fruit]
		if diff > 0 {
			for i := 0; i < diff/2; i++ {
				swapFrom1 = append(swapFrom1, fruit)
			}
		} else if diff < 0 {
			for i := 0; i < -diff/2; i++ {
				swapFrom2 = append(swapFrom2, fruit)
			}
		}
	}

	// Step 4: Sort extra fruits
	sort.Ints(swapFrom1)
	sort.Sort(sort.Reverse(sort.IntSlice(swapFrom2)))

	// Step 5: Find the smallest fruit value overall
	minFruit := math.MaxInt
	for _, x := range basket1 {
		minFruit = min(minFruit, x)
	}
	for _, x := range basket2 {
		minFruit = min(minFruit, x)
	}

	// Step 6: Calculate minimum total cost
	cost := 0
	for i := 0; i < len(swapFrom1); i++ {
		a := swapFrom1[i]
		b := swapFrom2[i]
		direct := min(a, b)
		indirect := 2 * minFruit
		cost += min(direct, indirect)
	}
	return int64(cost)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
