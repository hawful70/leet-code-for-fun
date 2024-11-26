func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    frequency := 0
    left := 0
    total := 0
    for right, num := range nums {
        total += num

        for nums[right] * (right - left + 1) - total > k {
            total -= nums[left]
            left += 1
        }

        frequency = max(frequency, right - left + 1)
    }
    return frequency
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}