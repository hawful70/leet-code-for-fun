class Solution:
    def nthUglyNumber(self, n: int, a: int, b: int, c: int) -> int:
        ab = self.lcm(a, b)
        ac = self.lcm(a, c)
        bc = self.lcm(b, c)
        abc = self.lcm(ab, c)

        left, right = 1, min(a, b, c) * n

        while left < right:
            mid = left + (right - left) // 2
            if self.count(mid, a, b, c, ab, ac, bc, abc) >= n:
                right = mid
            else:
                left = mid + 1

        return left

    def count(self, x: int, a: int, b: int, c: int,
              ab: int, ac: int, bc: int, abc: int) -> int:
        return (
            x // a + x // b + x // c
            - x // ab - x // ac - x // bc
            + x // abc
        )

    def gcd(self, a: int, b: int) -> int:
        while b != 0:
            a, b = b, a % b
        return a

    def lcm(self, a: int, b: int) -> int:
        return a // self.gcd(a, b) * b
