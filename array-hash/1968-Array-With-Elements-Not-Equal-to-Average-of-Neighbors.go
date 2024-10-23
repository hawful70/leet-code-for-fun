func rearrangeArray(nums []int) []int {
    for i := 1; i < len(nums) - 1; i++ {
        if (nums[i - 1] > nums[i] && nums[i] > nums[i + 1]) || (nums[i - 1] < nums[i] && nums[i] < nums[i + 1]) {
            nums[i], nums[i + 1] = nums[i + 1], nums[i]
        }
    }
    return nums
}