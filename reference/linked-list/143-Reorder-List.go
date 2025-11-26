/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first, second := head, prev
	for second != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next = second
		second.Next = temp1
		first, second = temp1, temp2
	}
}