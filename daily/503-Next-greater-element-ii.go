package daily

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stacks := []int{}
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	for i := 0; i < 2*n; i++ {
		currentNum := nums[i%n]

		for len(stacks) > 0 && nums[stacks[len(stacks)-1]] < currentNum {
			ans[stacks[len(stacks)-1]] = currentNum
			stacks = stacks[:len(stacks)-1]
		}
		if i < n {
			stacks = append(stacks, i)
		}
	}

	return ans
}
