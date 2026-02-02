class Solution:
    def isHappy(self, n: int) -> bool:
        seen = set()
        while n != 1 and n not in seen:
            seen.add(n)

            next_sum = 0
            while n > 0:
                digit = n % 10
                next_sum += digit ** 2
                n = n // 10

            n = next_sum 
        return n == 1