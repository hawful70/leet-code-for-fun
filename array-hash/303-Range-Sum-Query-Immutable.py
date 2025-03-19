class NumArray:

    def __init__(self, nums: List[int]):
        self.prefixNum = [0] * len(nums)
        self.prefixNum[0] = nums[0]

        for i in range(1, len(nums)):
            self.prefixNum[i] = self.prefixNum[i - 1] + nums[i]

    def sumRange(self, left: int, right: int) -> int:
        if left == 0:
            return self.prefixNum[right]
        return self.prefixNum[right] - self.prefixNum[left - 1]