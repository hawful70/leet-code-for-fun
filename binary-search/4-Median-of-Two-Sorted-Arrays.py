class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        
        len1, len2 = len(nums1), len(nums2)
        low, high = 0, len1

        while low <= high:
            partitionSmallArray = (low + high) // 2
            partitionLargeArray = (len1 + len2 + 1) // 2 - partitionSmallArray

            maxLeftSmall = float('-inf') if partitionSmallArray == 0 else nums1[partitionSmallArray - 1]
            minRightSmall = float('inf') if partitionSmallArray == len1 else nums1[partitionSmallArray]

            maxLeftLarge = float('-inf') if partitionLargeArray == 0 else nums2[partitionLargeArray - 1]
            minRightLarge = float('inf') if partitionLargeArray == len2 else nums2[partitionLargeArray]

            if maxLeftSmall <= minRightLarge and maxLeftLarge <= minRightSmall:
                if (len1 + len2) % 2 == 0:
                    return (max(maxLeftSmall, maxLeftLarge) + min(minRightSmall, minRightLarge)) / 2
                else:
                    return max(maxLeftSmall, maxLeftLarge)
            elif maxLeftSmall > minRightLarge:
                high = partitionSmallArray - 1
            else:
                low = partitionSmallArray + 1

            
