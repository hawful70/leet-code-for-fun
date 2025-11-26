func isPerfectSquare(num int) bool {
    left, right := 1, num
    for left <= right {
        mid := (left + right) / 2
        square := mid * mid
        if square == num {
            return true
        } else if square > num {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return false
}