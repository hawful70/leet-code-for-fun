class Solution:
    def suggestedProducts(self, products: List[str], searchWord: str) -> List[List[str]]:
        res = []
        products.sort()
        left, right = 0, len(products) - 1
        for i in range(len(searchWord)):
            s = searchWord[i]
            
            # eliminate unmatching word
            while left <= right and (len(products[left]) <= i or products[left][i] != s):
                left += 1
            while left <= right and (len(products[right]) <= i or products[right][i] != s):
                right -= 1

            res.append([])
            remain = right - left + 1
            for j in range(min(3, remain)):
                res[-1].append(products[left + j])
        return res