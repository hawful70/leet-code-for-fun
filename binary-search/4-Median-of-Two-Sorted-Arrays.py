class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        len1, len2 = len(nums1), len(nums2)

        low, high = 0, len1
        while low <= high:
            partitionSmallArr = (low + high) // 2
            partitionLargeArr = (len1 + len2 + 1) // 2 - partitionSmallArr

            maxLeftSmall = float('-inf') if partitionSmallArr == 0 else nums1[partitionSmallArr - 1]
            minRightSmall = float('inf') if partitionSmallArr == len1 else nums1[partitionSmallArr]
            
            maxLeftLarge = float('-inf') if partitionLargeArr == 0 else nums2[partitionLargeArr - 1]
            minRightLarge = float('inf') if partitionLargeArr == len2 else nums2[partitionLargeArr]

            if maxLeftSmall <= minRightLarge and maxLeftLarge <= minRightSmall:
                if (len1 + len2) % 2 == 0:
                    return (max(maxLeftSmall, maxLeftLarge) + min(minRightSmall, minRightLarge)) / 2
                else:
                    return max(maxLeftSmall, maxLeftLarge)
            elif maxLeftSmall > minRightLarge:
                high = partitionSmallArr - 1
            else:
                low = partitionSmallArr + 1
            
