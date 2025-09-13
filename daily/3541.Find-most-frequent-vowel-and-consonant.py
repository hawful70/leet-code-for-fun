class Solution:
    def maxFreqSum(self, s: str) -> int:
        vowels = set("aeiou")
        max_vowel = 0
        max_consonant = 0
        vcount = {}
        ccount = {}

        for ch in s:
            if ch in vowels:
                vcount[ch] = vcount.get(ch, 0) + 1
                if vcount[ch] > max_vowel:
                    max_vowel = vcount[ch]
            else:
                ccount[ch] = ccount.get(ch, 0) + 1
                if ccount[ch] > max_consonant:
                    max_consonant = ccount[ch]

        return max_vowel + max_consonant

