class Solution:
    def reverseWords(self, s: str) -> str:
        s_list = list(s)
        def util_reverword(l, h):
            while l < h:
                s_list[l], s_list[h] = s_list[h], s_list[l]
                l += 1
                h -= 1
            
        
        low = 0
        for high in range(1, len(s)):
            if s[high] == \ \:
                util_reverword(low, high - 1)
                low = high + 1
            elif high == len(s) - 1:
                util_reverword(low, len(s_list) - 1)
                
        return \\.join(s_list)
        


