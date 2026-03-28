class Solution:
    def mergeNodes(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None

        dummy = ListNode(0)
        curr = dummy

        zero_count = 0
        sum_nodes = 0
        while head:
            if head.val == 0:
                zero_count += 1 
            sum_nodes += head.val
            
            if zero_count >= 2:
                newNode = ListNode(sum_nodes) if sum_nodes > 0 else None
                if newNode:
                    curr.next = newNode
                    curr = curr.next
                zero_count = 1
                sum_nodes = 0

            head = head.next
        return dummy.next