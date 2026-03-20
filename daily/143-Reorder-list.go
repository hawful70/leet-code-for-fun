package daily

func reorderList(head *ListNode)  {
    nodes := []*ListNode{}
    curr := head
    for curr != nil {
        nodes = append(nodes, curr)
        curr = curr.Next
    }

    left, right := 0, len(nodes) - 1
    for left < right {
        nodes[left].Next = nodes[right]
        left++

        if left == right {
            break
        }

        nodes[right].Next = nodes[left]
        right--
    }
    nodes[left].Next = nil
}
