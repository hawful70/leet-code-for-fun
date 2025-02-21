class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        character_map = {}

        for char in s:
            character_map[char] = character_map.get(char, 0) + 1

        for char in t:
            character_map[char] = character_map.get(char, 0) - 1

        return all(value == 0 for value in character_map.values())