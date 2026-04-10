class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        if not l1 or not l2:
            return None

        l1_stacks, l2_stacks = [], []
        while l1:
            l1_stacks.append(l1.val)
            l1 = l1.next
        while l2:
            l2_stacks.append(l2.val)
            l2 = l2.next
        
        carry = 0
        head = None
        while l1_stacks or l2_stacks or carry:
            item1 = l1_stacks.pop() if l1_stacks else 0 
            item2 = l2_stacks.pop() if l2_stacks else 0
            total = item1 + item2 + carry
            carry = total // 10
            new_digit = total % 10
            
            newNode = ListNode(new_digit)
            newNode.next = head
            head = newNode

        return head