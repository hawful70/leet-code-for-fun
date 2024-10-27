func containsDuplicate(nums []int) bool {
    duplicate_map := make(map[int]int)
    for _, num := range nums {
        duplicate_map[num]++
        if duplicate_map[num] > 1 {
            return true
        }
    }
    return false
}