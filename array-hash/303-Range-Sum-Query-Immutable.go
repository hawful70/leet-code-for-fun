type NumArray struct {
	prefixNum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	return NumArray{
		prefixNum: append([]int{}, prefixSum...),
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixNum[right]
	}
	return this.prefixNum[right] - this.prefixNum[left-1]
}