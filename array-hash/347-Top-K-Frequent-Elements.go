func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	n := len(nums)
	frequency := make([][]int, n+1)

	for num, freq := range count {
		frequency[freq] = append(frequency[freq], num)
	}

	var res []int
	for i := n; i > 0; i-- {
		for _, num := range frequency[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}