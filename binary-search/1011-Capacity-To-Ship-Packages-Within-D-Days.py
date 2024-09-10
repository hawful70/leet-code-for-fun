class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        if len(weights) == 1:
            return weight[0]
        left, right = max(weights), sum(weights)
        ans = right
        while left <= right:
            mid = (left + right) // 2
            if self.validCapacity(weights, days, mid):
                ans = min(ans, mid)
                right = mid - 1
            else:
                left = mid + 1
        return ans
    def validCapacity(self, weights: List[int], days: int, capacity: int) -> bool:
        day_shipped = 1
        current_capacity = capacity
        for weight in weights:
            if weight > current_capacity:
                day_shipped += 1
                current_capacity = capacity
            current_capacity -= weight
        return day_shipped <= days