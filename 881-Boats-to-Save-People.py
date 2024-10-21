class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        if len(people) == 1 and people[0] <= limit:
            return 1
        if len(people) == 1 and people[0] > limit:
            return 0
        people.sort()
        ans = 0
        low, high = 0, len(people) - 1
        while low <= high:
            if people[low] + people[high] <= limit:
                low += 1
                high -= 1
            else:
                high -= 1
            ans += 1
        return ans              
