class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        if len(s) == 0:
            return 0
        s.sort()
        g.sort()
        
        g_pointer = len(g) - 1
        s_pointer = len(s) - 1
        ans = 0
        while s_pointer > -1 and g_pointer > -1:
            if g[g_pointer] <= s[s_pointer]:
                s_pointer -= 1
                ans += 1
            g_pointer -= 1
        return ans