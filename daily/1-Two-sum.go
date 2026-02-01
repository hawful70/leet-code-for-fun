package daily

func twoSum(nums []int, target int) []int {
	mapSum := make(map[int]int)
	for index, num := range nums {
		remain := target - num
		if value, ok := mapSum[remain]; ok {
			return []int{index, value}
		} else {
			mapSum[num] = index
		}
	}
	return []int{}
}
