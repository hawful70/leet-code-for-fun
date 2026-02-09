package daily

func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    left := 0
    current_sum := 0
    res := math.MaxInt32
    for right := 0; right < n; right++ {
        current_sum += nums[right]
        for current_sum <= target {
            res = min(res, right - left + 1)
            current_sum -= nums[left]
            left++
        }
    }
    
    if res == math.MaxInt32 {
        return 0
    }
    return res
}