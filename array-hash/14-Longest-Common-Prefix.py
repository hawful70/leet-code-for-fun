class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""

        prefix = strs[0]
        for i in range(1, len(strs)):
            while len(strs[i]) < len(prefix):
                prefix = prefix[:len(prefix) - 1]

            while strs[i][:len(prefix)] != prefix:
                prefix = prefix[:len(prefix) - 1]
                if not prefix:
                    return ""

        return prefix
