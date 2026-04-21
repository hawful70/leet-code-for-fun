class Solution:
    def splitListToParts(self, head: Optional[ListNode], k: int) -> List[Optional[ListNode]]:
        ans = [None] * k
        size = 0
        current = head
        while current:
            size += 1
            current = current.next
        
        splitSize = size // k
        remainSize = size % k
        current = head
        for i in range(k):
            ans[i] = current
            currentSize = splitSize + (1 if remainSize > 0 else 0)
            remainSize -= 1

            for _ in range(currentSize - 1):
                if current:
                    current = current.next
            if current:
                nxt = current.next
                current.next = None
                current = nxt

        return ans