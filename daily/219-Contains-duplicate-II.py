class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        seen = set()
        n = len(nums)
        left, right = 0, 0
        for right in range(n):
            while nums[right] in seen:
                if abs(right - left) > k:
                    seen.remove(nums[left])
                    left += 1
                else:
                    return True
            seen.add(nums[right])

        return False 
            
