func numSubseq(nums []int, target int) int {
    const mod = 1000000007
    sort.Ints(nums)

    ans := 0
    left, right := 0, len(nums) - 1
    for left <= right {
        sum := nums[left] + nums[right]
        if sum <= target {
            ans += modPow(2, right - left, mod)
            ans = ans % mod
            left++
        } else {
            right--
        }
    }
    return ans
}

func modPow(base, exp, mod int) int {
    result := 1
    base = base % mod
    for exp > 0 {
        if exp % 2 == 1 { // If exp is odd
            result = (result * base) % mod
        }
        base = (base * base) % mod // Square the base
        exp = exp / 2
    }
    return result
}
