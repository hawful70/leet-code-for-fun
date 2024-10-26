func trap(height []int) int {
    ans := 0
    left, right := 0, len(height) - 1
    maxLeft, maxRight := height[left], height[right]
    for left < right {
        if maxLeft < maxRight {
            left++
            maxLeft = max(maxLeft, height[left])
            ans += maxLeft - height[left]
        } else {
            right--
            maxRight = max(maxRight, height[right])
            ans += maxRight - height[right]
        }
    }
    return ans
}

func max(a int, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}