func minSubArrayLen(target int, nums []int) int {
    min_len := len(nums)

    current_sum := 0
    left := 0
    isFound := false
    for right, _ := range nums {
        current_sum += nums[right]

        for current_sum >= target {
            min_len = min(min_len, right - left + 1)
            current_sum -= nums[left]
            isFound = true
            left++
        }
    }

    if isFound {
        return min_len
    } else {
        return 0
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}