class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        minDistance, currentDistance = 0, 0
        for i in range(k):
            currentDistance += abs(arr[i] - x)
        
        startIndex = 0
        minDistance = currentDistance
        for i in range(1, len(arr) - k + 1):
            currentDistance += abs(arr[i+k-1] - x) - abs(arr[i-1] - x)
            if currentDistance < minDistance:
                minDistance = currentDistance
                startIndex = i
        
        return arr[startIndex: startIndex+k]