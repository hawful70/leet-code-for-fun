func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	frequencyNumber := make(map[int]int)
	for _, num := range nums {
		if _, exists := frequencyNumber[num]; !exists {
			frequencyNumber[num] = 1
		} else {
			frequencyNumber[num]++
		}
	}

	n := len(nums)
	for key, value := range frequencyNumber {
		if value > (n / 2) {
			return key
		}
	}
	return 0
}