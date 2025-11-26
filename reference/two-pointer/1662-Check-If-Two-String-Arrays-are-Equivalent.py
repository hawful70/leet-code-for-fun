class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        # str1 = ""
        # str2 = ""
        # for c in word1:
        #     str1 += c
        # for c in word2:
        #     str2 += c
        # return str1 == str2
        i = j = 0
        pos1 = pos2 = 0
        while i < len(word1) and j < len(word2):
            if word1[i][pos1] != word2[j][pos2]:
                return False
            
            pos1 += 1
            pos2 += 1

            if pos1 == len(word1[i]):
                i += 1
                pos1 = 0
            if pos2 == len(word2[j]):
                j += 1
                pos2 = 0
        return i == len(word1) and j == len(word2)