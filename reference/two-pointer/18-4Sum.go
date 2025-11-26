func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    var results [][]int
    var result []int
    sum_up(4, 0, len(nums) - 1, nums, target, result, &results)
    return results
}

func sum_up(N int, left int, right int, nums []int, target int, result []int, results *[][]int) {
    if N < 2 || right - left + 1 < N || target < nums[left] * N || target > nums[right] * N {
        return
    }

    if N == 2 {
        for left < right {
            current_sum := nums[left] + nums[right]
            if current_sum == target {
                temp := make([]int, len(result))
                copy(temp, result)

                temp = append(temp, nums[left], nums[right])
                *results = append(*results, temp)

                left++
                for left < right && nums[left] == nums[left - 1] {
                    left++
                }
            } else if current_sum > target {
                right--
            } else {
                left++
            }
        }
    } else {
        for i := left; i < len(nums); i++ {
            if i > left && nums[i] == nums[i - 1] {
                continue
            }

            result = append(result, nums[i])
            sum_up(N - 1, i + 1, right, nums, target - nums[i], result, results)
            result = result[:len(result) - 1]
        }
    }
}