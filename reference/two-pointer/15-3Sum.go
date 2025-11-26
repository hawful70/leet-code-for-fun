func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var result [][]int

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        left, right := i + 1, len(nums) - 1
        for left < right {
            current_sum := nums[i] + nums[left] + nums[right]
            if current_sum > 0 {
                right -= 1
            } else if current_sum < 0 {
                left += 1
            } else {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                left += 1
                right -= 1
                for left < right && nums[left] == nums[left - 1] {
                    left += 1
                }
            }
        }
    }
    return result
}