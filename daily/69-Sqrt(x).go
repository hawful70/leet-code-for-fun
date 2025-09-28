func mySqrt(x int) int {
    left, right := 1, x
    ans := 0
    for left <= right {
        mid := (left + right) / 2
        square := mid * mid
        if square == x {
            return mid
        } else if square < x {
            ans = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return ans
}