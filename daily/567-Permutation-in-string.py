class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        target_window = Counter(s1)
        k = len(s1)
        current_window = defaultdict(int)
        for i in range(k):
            current_window[s2[i]] += 1
        
        for i in range(k, len(s2)):
            if target_window == current_window:
                return True
            current_window[s2[i]] += 1
            current_window[s2[i - k]] -= 1
            if current_window[s2[i - k]] == 0:
                del current_window[s2[i - k]]

        return target_window == current_window
