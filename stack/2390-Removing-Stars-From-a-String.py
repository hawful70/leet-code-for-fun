class Solution:
    def removeStars(self, s: str) -> str:
        stacks = []
        for c in s:
            if stacks and c == "*" and stacks[-1] != "*":
                stacks.pop()
            else:    
                stacks.append(c)
        return "".join(stacks)