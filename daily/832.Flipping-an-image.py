from typing import List


class Solution:
    def flipAndInvertImage(self, image: List[List[int]]) -> List[List[int]]:
        rows = len(image)
        cols = len(image[0]) if rows else 0
        for i in range(rows):
            left, right = 0, cols - 1
            while left <= right:
                if image[i][left] == image[i][right]:
                    image[i][left] ^= 1
                    image[i][right] = image[i][left]
                left += 1
                right -= 1
        return image

    def flipAndInvertImageV2(self, image: List[List[int]]) -> List[List[int]]:
        rows = len(image)
        cols = len(image[0]) if rows else 0
        for i in range(rows):
            left, right = 0, cols - 1
            while left <= right:
                image[i][left], image[i][right] = image[i][right], image[i][left]
                left += 1
                right -= 1
        for i in range(rows):
            for j in range(cols):
                image[i][j] ^= 1
        return image

