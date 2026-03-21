class Solution:
    def nextLargerNodes(self, head: Optional[ListNode]) -> List[int]:
        stacks = []
        greater_node = {}
        curr = head
        while curr:
            greater_node[curr] = 0
            while stacks and stacks[-1].val < curr.val:
                greater_node[stacks[-1]] = curr.val
                stacks.pop()
            stacks.append(curr)
            curr = curr.next
        
        return list(greater_node.values())