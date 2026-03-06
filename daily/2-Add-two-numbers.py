class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        curr = dummy
        curry = 0
        while l1 or l2 or curry > 0:
            if l1:
                curry += l1.val
                l1 = l1.next
            if l2:
                curry += l2.val
                l2 = l2.next

            curr.next = ListNode(curry % 10)
            curr = curr.next
            curry //= 10
        return dummy.next