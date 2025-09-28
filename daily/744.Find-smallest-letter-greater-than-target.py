from typing import List


class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        left, right = 0, len(letters) - 1
        candidate = letters[0]
        found = False

        while left <= right:
            mid = (left + right) // 2
            if letters[mid] > target:
                candidate = letters[mid]
                right = mid - 1
                found = True
            else:
                left = mid + 1

        if found:
            return candidate
        return letters[0]
