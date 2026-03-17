class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        dummy = ListNode(0, list1)
        before = dummy

        for _ in range(a):
            before = before.next
        
        curr = before.next
        for _ in range(b - a + 1):
            curr = curr.next
        
        list2_curr = list2
        while list2_curr.next:
            list2_curr = list2_curr.next
            
        list2_curr.next = curr
        before.next = list2
        return dummy.next