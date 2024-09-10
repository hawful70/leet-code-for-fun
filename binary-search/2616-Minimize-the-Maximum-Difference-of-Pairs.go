
import "sort"
import "math"
func minimizeMax(nums []int, p int) int {
    if len(nums) == 1 {
        return 0
    }

    sort.Ints(nums)
    left, right := 0, int(int64(math.Pow10(9)))
    ans := 0
    for left <= right {
        mid := (left + right) / 2
        if validMininumDifference(nums, p, mid) {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

func validMininumDifference(nums []int, p int, threehold int) bool {
    i, count := 0, 0
    for i < len(nums) - 1 {
        diff := int(math.Abs(float64(nums[i] - nums[i+1])))
        if diff <= threehold {
            count += 1
            i += 2
        } else {
            i += 1
        }

        if count == p {
            return true
        }
    }
    return false
}