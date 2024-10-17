func maxArea(height []int) int {
    ans := 0
    left, right := 0, len(height) - 1
    for left < right {
        minHeight := min(height[left], height[right])
        currentContainer := minHeight * (right - left)
        if currentContainer > ans {
            ans = currentContainer
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return ans
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}