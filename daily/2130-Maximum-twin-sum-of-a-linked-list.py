class Solution:
    def pairSum(self, head: Optional[ListNode]) -> int:
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        prev = None
        curr = slow
        while curr:
            next_node = curr.next
            curr.next = prev
            prev = curr
            curr = next_node
        
        maxTwinSum = 0
        left, right = head, prev
        while right:
            currentTwinSum = left.val + right.val
            maxTwinSum = max(maxTwinSum, currentTwinSum)
            left = left.next
            right = right.next

        return maxTwinSum