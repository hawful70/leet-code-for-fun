func fourSum(nums []int, target int) [][]int {
    var results [][]int
    var result []int

    sort.Ints(nums)
    backtrack(nums, 4, target, 0, len(nums) - 1, result, &results)
    return results
}

func backtrack(nums []int, N int, target int, left int, right int, result []int, results *[][]int) {
    if right - left + 1 < N || N < 2 || target < nums[left] * N || target > nums[right] * N {
        return
    }

    if N == 2 {
        for left < right {
            s := nums[left] + nums[right]
            if s == target {
                temp := make([]int, len(result))
                copy(temp, result)

                temp = append(temp, nums[left], nums[right])
                *results = append(*results, temp)
                left++
                for left < right && nums[left] == nums[left - 1] {
                    left++
                }
            } else if s > target {
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
            backtrack(nums, N - 1, target - nums[i], i + 1, len(nums) - 1, result, results)
            result = result[:len(result) - 1]
        } 
    }
}

