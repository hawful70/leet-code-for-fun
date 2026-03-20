class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None
        old_node = {}
        curr = head
        while curr:
            old_node[curr] = Node(curr.val)
            curr = curr.next
        
        curr = head
        while curr:
            old_node[curr].next = old_node.get(curr.next)
            old_node[curr].random = old_node.get(curr.random)
            curr = curr.next
        return old_node[head]