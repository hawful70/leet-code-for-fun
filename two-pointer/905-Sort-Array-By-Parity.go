func sortArrayByParity(nums []int) []int {
    if len(nums) == 1 {
        return nums
    }

    left, right := 0, 1
    for left <= right && right < len(nums) {
        if nums[left] % 2 == 0 {
            left++
        }
        if nums[left] % 2 != 0 && nums[right] % 2 == 0 {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }
        right++
    }
    return nums
}