package daily

func containsNearbyDuplicate(nums []int, k int) bool {
    n := len(nums)
    seen := make(map[int]bool)
    left, right := 0, 0
    for right < n {
        for seen[nums[right]] {
            if abs(right - left) > k {
                delete(seen, nums[left])
                left++
            } else {
                return true
            }
        }
        seen[nums[right]] = true
        right++
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}