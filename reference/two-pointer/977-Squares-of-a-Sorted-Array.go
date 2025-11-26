func sortedSquares(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    left, right := 0, n - 1
    index := n - 1

    for left <= right {
        square_left := nums[left] * nums[left]
        square_right := nums[right] * nums[right]

        if square_left > square_right {
            // ans[right - left] = square_left
            ans[index] = square_left
            left++
        } else {
            // ans[right - left] = square_right
            ans[index] = square_right
            right--
        }

        index --
    }
    return ans
}