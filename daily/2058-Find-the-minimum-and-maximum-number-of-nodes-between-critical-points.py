class Solution:
    def nodesBetweenCriticalPoints(self, head: Optional[ListNode]) -> List[int]:
        if not head or not head.next or not head.next.next:
            return [-1, -1]
        
        first_critical_point, last_critical_point, minDistance = -1, -1, float('inf')
        prev, curr = head, head.next
        current_index = 1
        while curr.next:
            next_node = curr.next
            maxima = curr.val > prev.val and curr.val > next_node.val
            minima = curr.val < prev.val and curr.val < next_node.val
            if maxima or minima:
                if first_critical_point == -1:
                    first_critical_point = current_index
                else:
                    minDistance = min(minDistance, current_index - last_critical_point)
                last_critical_point = current_index
            prev = curr
            curr = next_node
            current_index += 1
        
        if first_critical_point == last_critical_point:
            return [-1, -1]
        return [minDistance, last_critical_point - first_critical_point]