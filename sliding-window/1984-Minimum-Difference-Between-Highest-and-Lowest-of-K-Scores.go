func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    right := k - 1
    ans := math.MaxInt32

    for left := 0; left < len(nums); left++ {
        if right < len(nums) {
            ans = min(nums[right] - nums[left], ans)
            right++
        }
    }
    return ans
}