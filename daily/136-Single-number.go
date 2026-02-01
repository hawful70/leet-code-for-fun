package daily

func singleNumber(nums []int) int {
	map_num := make(map[int]int)
	for _, num := range nums {
		map_num[num]++
	}

	for key, value := range map_num {
		if value != 2 {
			return key
		}
	}
	return 0
}
