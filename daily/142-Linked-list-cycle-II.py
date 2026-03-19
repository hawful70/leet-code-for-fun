class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # visitNode = set()
        # while head:
        #     if head in visitNode:
        #         return head
        #     visitNode.add(head)
        #     head = head.next
        # return None
        fast = slow = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                break
        else:
            return None
        
        fast = head
        while fast != slow:
            fast = fast.next
            slow = slow.next
        return fast