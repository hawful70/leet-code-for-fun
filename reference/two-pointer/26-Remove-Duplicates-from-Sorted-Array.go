func removeDuplicates(nums []int) int {
    found := 0
    for i := 1; i < len(nums); i++ {
        if nums[found] != nums[i] {
            found++
            nums[found], nums[i] = nums[i], nums[found]
        }
    }
    return found + 1
}