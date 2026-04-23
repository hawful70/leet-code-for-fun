class Solution:
    def insertionSortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)   # sorted list starts empty
        curr = head

        while curr:
            next_node = curr.next   # save next before relinking
            prev = dummy

            # Find the correct position in sorted portion
            while prev.next and prev.next.val < curr.val:
                prev = prev.next

            # Splice curr into sorted portion
            curr.next = prev.next
            prev.next = curr

            curr = next_node  # move to next unsorted node

        return dummy.next