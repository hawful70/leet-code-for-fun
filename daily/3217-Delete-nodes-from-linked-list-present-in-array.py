class Solution:
    def modifiedList(self, nums: List[int], head: Optional[ListNode]) -> Optional[ListNode]:
        seen = set()
        for num in nums:
            seen.add(num)
        
        dummy = ListNode(0)
        curr = dummy
        while head:
            if head.val not in seen:
                curr.next = head
                curr = curr.next
            head = head.next
        curr.next = None
        return dummy.next