func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    total := 0
    frequency := 0

    left := 0
    for right, _ := range nums {
        total += nums[right]

        for nums[right] * (right - left + 1) - total > k {
            total -= nums[left]
            left++
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