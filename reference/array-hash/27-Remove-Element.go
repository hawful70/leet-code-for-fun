func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == val {
			if nums[left] != nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
			right--
		} else {
			left++
		}
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			k++
		} else {
			break
		}
	}
	return k
}

// another solution
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}