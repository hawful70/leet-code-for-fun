class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagram_hash = defaultdict(list)

        for s in strs:
            ordered_str = "".join(sorted(s))  # Sort the string to get the key
            anagram_hash[ordered_str].append(s)

        return list(anagram_hash.values())