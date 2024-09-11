class Solution:
    def maximumRemovals(self, s: str, p: str, removable: List[int]) -> int:
        left, right = 0, len(removable)
        ans = 0
        while left <= right:
            mid = (left + right) // 2
            if self.isValidRemoval(mid, s, p, removable):
                ans = max(mid, ans)
                left = mid + 1
            else:
                right = mid - 1
        return ans

    def isValidRemoval(self, subsequence: int, s: str, p: str, removable: List[int]) -> bool:
        removal = set(removable[:subsequence])
        p_pointer = 0
        for i in range(len(s)):
            if i in removal:
                continue
            if s[i] == p[p_pointer]:
                p_pointer += 1
                if p_pointer == len(p):
                    return True
        return False

