class Solution:
    def firstPalindrome(self, words: List[str]) -> str:
        def util_simultaneously(word):
            low, high = 0, len(word) - 1
            while low < high:
                if word[low] != word[high]:
                    return False
                low += 1
                high -= 1
            return True
            
        for word in words:
            if util_simultaneously(word):
                return word
        return ""