func numOfSubarrays2(arr []int, k int, threshold int) int {
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

func numOfSubarrays(arr []int, k int, threshold int) int {
    ans := 0
    current_sum := sum_arr(0, k - 1, arr)
    for i := 0; i < len(arr) - k + 1; i++ {
        current_sum += arr[i + k - 1]
        current_average := float64(current_sum) / float64(k)
        if current_average >= float64(threshold) {
            ans++
        }
        current_sum -= arr[i]
    }
    return ans
}

func sum_arr(left int, right int, arr []int) int {
    sum := 0
    for i := left; i < right; i++ {
        sum += arr[i]
    }
    return sum
}