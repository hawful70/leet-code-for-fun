from typing import List


class Solution:
    def longestMountain(self, arr: List[int]) -> int:
        n = len(arr)
        if n < 3:
            return 0

        ans = 0
        left = 0
        while left + 2 < n:
            right = left

            if arr[right] < arr[right + 1]:
                while right + 1 < n and arr[right] < arr[right + 1]:
                    right += 1

                if right + 1 < n and arr[right] > arr[right + 1]:
                    while right + 1 < n and arr[right] > arr[right + 1]:
                        right += 1
                    ans = max(ans, right - left + 1)

            left = right if right > left else left + 1

        return ans

