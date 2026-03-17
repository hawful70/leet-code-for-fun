class Solution:
    def removeNodes(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None
        stack = deque()
        curr = head
        while curr:
            while stack and stack[-1].val < curr.val:
                stack.pop()
            stack.append(curr)
            curr = curr.next
        
        new_head = stack[0]
        curr_node = new_head
        for i in range(1, len(stack)):
            curr_node.next = stack[i]
            curr_node = curr_node.next
        curr_node.next = None
        return new_head