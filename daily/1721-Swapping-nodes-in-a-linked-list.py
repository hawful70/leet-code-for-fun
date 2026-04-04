class Solution:
    def swapNodes(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        left = head
        for _ in range(k - 1):
            left = left.next
        
        fast = left
        slow = head
        while fast.next:
            slow = slow.next
            fast = fast.next
        
        left.val, slow.val = slow.val, left.val
        return head