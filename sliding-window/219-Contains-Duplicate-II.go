func containsNearbyDuplicate(nums []int, k int) bool {
    duplicate_map := make(map[int]bool)
    left := 0
    for right, num := range nums {
        for duplicate_map[num] {
            if right - left > k {
                delete(duplicate_map, nums[left])
                left++
            } else {
                return true
            }
        } 
        duplicate_map[num] = true
    }
    return false
}