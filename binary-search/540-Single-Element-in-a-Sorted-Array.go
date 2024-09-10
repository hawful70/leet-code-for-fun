func singleNonDuplicate(nums []int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := (left + right) / 2
        half_even := (right - mid) % 2 == 0
        if nums[mid + 1] == nums[mid] {
            if half_even {
                left = mid + 2
            } else {
                right = mid - 1
            }
        } else if nums[mid - 1] == nums[mid] {
            if half_even {
                right = mid - 2
            } else {
                left = mid + 1
            }
        } else {
            return nums[mid]
        }
    }
    return nums[left]
}