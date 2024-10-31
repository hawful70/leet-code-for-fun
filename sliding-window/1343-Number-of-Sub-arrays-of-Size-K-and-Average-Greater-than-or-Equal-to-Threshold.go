func numOfSubarrays(arr []int, k int, threshold int) int {
    ans := 0
    left := 0
    for i := k - 1; i < len(arr); i++ {
        current_sum := sum_arr(left, i, arr)
        current_average := float64(current_sum) / float64(k)
        if current_average >= float64(threshold) {
            ans++
        }
        left++
    }
    return ans
}

func sum_arr(left int, right int, arr []int) int {
    sum := 0
    for i := left; i <= right; i++ {
        sum += arr[i]
    }
    return sum
}