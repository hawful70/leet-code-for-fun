class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        seen = set()

        curr_A = headA
        while curr_A:
            seen.add(curr_A)
            curr_A = curr_A.next
        
        curr_B = headB
        while curr_B:
            if curr_B in seen:
                return curr_B
            seen.add(curr_B)
            curr_B = curr_B.next
        return None