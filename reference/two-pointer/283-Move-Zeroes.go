func moveZeroes(nums []int)  {
    if len(nums) == 1 {
        return
    }

    found := 0
    for i := 1; i < len(nums); i++ {
        if nums[found] == 0 && nums[i] != 0 {
            nums[found], nums[i] = nums[i], nums[found]
            found++
        }
        if nums[found] != 0 && nums[i] == 0 {
            found = i
        }
    }
}