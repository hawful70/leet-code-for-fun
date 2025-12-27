class Solution:
    def findPeakGrid(self, mat: List[List[int]]) -> List[int]:
        rows, cols = len(mat), len(mat[0])
        lo, hi = 0, rows - 1
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            curCol = 0
            for c in range(1, cols):
                if mat[mid][c] > mat[mid][curCol]:
                    curCol = c
            
            curEle = mat[mid][curCol]
            up, down = 0, 0
            if mid > 0:
                up = mat[mid - 1][curCol]
            if mid < rows - 1:
                down = mat[mid + 1][curCol]
            
            if mid > 0 and up > curEle:
                hi = mid - 1
            elif mid < rows -1 and down > curEle:
                lo = mid + 1
            else:
                return [mid, curCol]

        return [-1, -1]