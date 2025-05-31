/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle2(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	current := head
	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
