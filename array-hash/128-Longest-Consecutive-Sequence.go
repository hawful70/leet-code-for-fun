func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	maxlen := 0
	for num := range set {
		if !set[num-1] {
			length := 1
			current := num

			for set[current+1] {
				current += 1
				length++
			}

			if length > maxlen {
				maxlen = length
			}
		}
	}
	return maxlen
}