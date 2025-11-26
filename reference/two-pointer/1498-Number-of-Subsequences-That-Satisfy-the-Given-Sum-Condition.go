func numSubseq(nums []int, target int) int {
    const mod = 1000000007
    ans := 0
    sort.Ints(nums)
    left, right := 0, len(nums) - 1
    for left <= right {
        if nums[left] + nums[right] <= target {
            ans += powMod(2, right - left, mod)
            ans %= mod
            left++
        } else {
            right--
        }
    }
    return ans
}

func powMod(base, exp, mod int) int {
    result := 1
    base = base % mod
    for exp > 0 {
        if exp % 2 == 1 {
            result = (result * base) % mod
        }
        exp = exp >> 1
        base = (base * base) % mod
    }
    return result
}