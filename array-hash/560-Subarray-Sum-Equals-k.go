func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	countMap := make(map[int]int)
	countMap[0] = 1

	for _, num := range nums {
		prefixSum += num
		if val, found := countMap[prefixSum-k]; found {
			count += val
		}

		countMap[prefixSum]++
	}

	return count
}