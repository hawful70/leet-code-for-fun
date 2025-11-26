func numOfSubarrays(arr []int, k int, threshold int) int {
    ans := 0
    cur_sum := 0
    for i := 0; i < k - 1; i++ {
        cur_sum += arr[i]
    }
    for i := 0; i <= len(arr) - k; i++ {
        cur_sum += arr[i + k - 1]
        if (cur_sum / k) >= threshold {
            ans += 1
        }
        cur_sum -= arr[i]
    }
    return ans
}