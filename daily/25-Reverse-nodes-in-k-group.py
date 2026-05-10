class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None
        length = 0
        curr = head
        while curr:
            length += 1
            curr = curr.next
        
        groups = length // k
        dummy = ListNode(0, head)
        group_prev = dummy
        curr = head
        for _ in range(groups):
            group_detail = curr
            prev = None
            for _ in range(k):
                if curr:
                    nxt = curr.next
                    curr.next = prev
                    prev = curr
                    curr = nxt
            group_prev.next = prev
            group_detail.next = curr
            group_prev = group_detail
        return dummy.next
            