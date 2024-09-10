func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    left, right := 0, len(nums) - 1
    mid := 0
    for left <= right {
        mid = (left + right) / 2
        if (mid >= 1 && mid == right && nums[mid] > nums[mid - 1]) || (mid == left && nums[mid] > nums[mid + 1]) || (nums[mid] > nums[mid + 1] && nums[mid] > nums[mid - 1]) {
            return mid
        } else if nums[mid] < nums[mid + 1] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return mid
}