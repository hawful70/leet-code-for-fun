func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if idx, found := hashMap[complement]; found {
            return []int{idx, i}
        }
        hashMap[num] = i
    }

    return nil
}