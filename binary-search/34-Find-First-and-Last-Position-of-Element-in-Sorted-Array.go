func searchRange(nums []int, target int) []int {
    if len(nums) == 1 && nums[0] == target {
        return []int{0, 0}
    }
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            if nums[left] != nums[mid] {
                left += 1
            } else if nums[right] != nums[mid] {
                right -= 1
            } else {
                return []int{left, right}
            }
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return []int{-1, -1}
}