class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        product = [1] * len(nums)  # Initialize product array with 1s
        prefix_left, prefix_right = 1, 1
        
        # Left pass: Compute prefix product from the left
        for i in range(len(nums)):
            if i != 0:
                prefix_left *= nums[i - 1]
            product[i] = prefix_left
        
        # Right pass: Compute prefix product from the right
        for i in range(len(nums) - 1, -1, -1):
            if i != len(nums) - 1:
                prefix_right *= nums[i + 1]
            product[i] *= prefix_right
        
        return product