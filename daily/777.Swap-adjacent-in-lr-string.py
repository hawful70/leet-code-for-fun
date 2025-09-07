class Solution:
    def canTransform(self, start: str, end: str) -> bool:
        i = j = 0
        n = len(start)

        while i < n and j < n:
            # Skip 'X' in both strings
            while i < n and start[i] == 'X':
                i += 1
            while j < n and end[j] == 'X':
                j += 1

            if i == n or j == n:
                break

            if start[i] != end[j]:
                return False

            # 'L' can only move left: its index cannot increase
            if start[i] == 'L' and i < j:
                return False
            # 'R' can only move right: its index cannot decrease
            if start[i] == 'R' and i > j:
                return False

            i += 1
            j += 1

        # Remaining characters must be 'X' in both strings
        while i < n:
            if start[i] != 'X':
                return False
            i += 1
        while j < n:
            if end[j] != 'X':
                return False
            j += 1

        return True

