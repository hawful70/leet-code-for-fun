/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return \t     -1 if num is higher than the picked number
 *\t\t\t      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left, right := 1, n
    ans := 0
    for left <= right {
        mid := (left + right) / 2
        if guess(mid) == -1 {
            right = mid - 1
        } else if guess(mid) == 1 {
            left = mid + 1
        } else {
            ans = mid
            return ans
        }
    }

    return ans
}