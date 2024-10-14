func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers) - 1
    for left < right {
        current_sum := numbers[left] + numbers[right]
        if current_sum == target {
            return []int{left + 1, right + 1}
        } else if current_sum > target {
            right -= 1
        } else {
            left += 1
        }
    }
    return []int{}
}