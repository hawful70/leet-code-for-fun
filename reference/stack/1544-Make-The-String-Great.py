class Solution:
    def makeGood(self, s: str) -> str:
        stacks = []

        for c in s:
            if stacks and stacks[-1] != c and stacks[-1].lower() == c.lower():
                stacks.pop()
            else:
                stacks.append(c)
        return "".join(stacks)