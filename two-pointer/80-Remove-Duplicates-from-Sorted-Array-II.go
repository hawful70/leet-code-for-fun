func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }

    found := 2
    for i := 2; i < len(nums); i++ {
        if nums[i] != nums[found - 2] {
            nums[i], nums[found] = nums[found], nums[i]
            found++
        }
    }
    return found
}