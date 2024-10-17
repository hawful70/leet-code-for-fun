func fourSum(nums []int, target int) [][]int {
    var results [][]int
    var result []int

    sort.Ints(nums)
    backtrack(nums, 0, target, 4, result, &results)
    return results
}

func backtrack(nums []int, start, target, N int, result []int, results *[][]int) {
    if N < 2 || start >= len(nums) || nums[start] * N > target || target > nums[len(nums) - 1] * N {
        return
    }

    if N == 2 {
        left, right := start, len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                temp := make([]int, len(result))
                copy(temp, result)

                temp = append(temp, nums[left], nums[right])
                *results = append(*results, temp)
                left++

                for left < right && nums[left] == nums[left - 1] {
                    left++
                }
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
        return
    }

    for i := start; i < len(nums) - N + 1; i++ {
        if i > start && nums[i] == nums[i - 1] {
            continue
        }

        result = append(result, nums[i])
        backtrack(nums, i + 1, target - nums[i], N - 1, result, results)
        result = result[:len(result) - 1]
    }
}