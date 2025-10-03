from typing import List


class Solution:
    def bagOfTokensScore(self, tokens: List[int], power: int) -> int:
        tokens.sort()
        score = 0
        best = 0
        left, right = 0, len(tokens) - 1

        while left <= right and (power >= tokens[left] or score > 0):
            if power >= tokens[left]:
                power -= tokens[left]
                score += 1
                best = max(best, score)
                left += 1
            elif score > 0:
                power += tokens[right]
                score -= 1
                right -= 1

        return best
