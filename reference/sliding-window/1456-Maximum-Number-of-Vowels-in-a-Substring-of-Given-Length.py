class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowel_letters = {"a", "e", "i", "o", "u"}  

        current_vowel = 0
        max_number, left = 0, 0
        for right in range(len(s)):
            if s[right] in vowel_letters:
                current_vowel += 1
            if right - left + 1 > k:
                if s[left] in vowel_letters:
                    current_vowel -= 1
                
                left += 1
            
            max_number = max(max_number, current_vowel)
        return max_number

