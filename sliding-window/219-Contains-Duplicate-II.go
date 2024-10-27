func containsNearbyDuplicate(nums []int, k int) bool {
    window := make(map[int]bool)
    left := 0
    for right, num := range nums {
        if right - left > k {
            delete(window, nums[left])
            left++
        }
        if window[num] {
            return true
        }
        window[num] = true
    }
    return false
}