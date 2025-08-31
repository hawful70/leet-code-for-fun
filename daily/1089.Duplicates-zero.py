from typing import List


class Solution:
    def duplicateZeros(self, arr: List[int]) -> None:
        n = len(arr)
        zeros = 0
        for x in arr:
            if x == 0:
                zeros += 1
        i, j = n - 1, n + zeros - 1
        while i < j:
            if arr[i] != 0:
                if j < n:
                    arr[j] = arr[i]
            else:
                if j < n:
                    arr[j] = 0
                j -= 1
                if j < n:
                    arr[j] = 0
            i -= 1
            j -= 1

    def duplicateZerosV2(self, arr: List[int]) -> None:
        n = len(arr)
        i = 0
        while i < n:
            if arr[i] == 0 and i < n - 1:
                for j in range(n - 2, i - 1, -1):
                    arr[j + 1] = arr[j]
                i += 1
            i += 1

