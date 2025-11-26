class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        word1Index, word2Index = 0, 0
        ans = ""
        while word1Index < len(word1) and word2Index < len(word2):
            ans += word1[word1Index]
            ans += word2[word2Index]
            word1Index += 1
            word2Index += 1
        
        while word1Index < len(word1):
            ans += word1[word1Index]
            word1Index += 1
        while word2Index < len(word2):
            ans += word2[word2Index]
            word2Index += 1
        return ans
        