class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        map_s = {}
        seen_t = set()

        for i in range(len(s)):
            char_s = s[i]
            char_t = t[i]

            if char_s in map_s:
                if map_s[char_s] != char_t:
                    return False
            else:
                if char_t in seen_t:
                    return False

                map_s[char_s] = char_t
                seen_t.add(char_t)
        return True