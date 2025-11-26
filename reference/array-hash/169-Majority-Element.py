class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]

        frequency_number = {}
        for num in nums:
            if num not in frequency_number:
                frequency_number[num] = 1
            else:
                frequency_number[num] += 1

        n = len(nums)
        for key, value in frequency_number.items():
            if value > n // 2:
                return key

        return 0