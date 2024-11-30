class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        trees = {}
        left, max_fruits = 0, 0

        for right, fruit in enumerate(fruits):
            if fruit not in trees:
                trees[fruit] = 1
            else:
                trees[fruit] += 1

            while len(trees) > 2:
                trees[fruits[left]] -= 1
                if trees[fruits[left]] == 0:
                    del trees[fruits[left]]
                left += 1

            max_fruits = max(max_fruits, right - left + 1)

        return max_fruits