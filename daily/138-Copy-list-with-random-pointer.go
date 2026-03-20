package daily

type Node struct {
    Val int
    Next *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    
    oldToNew := make(map[*Node]*Node)
    
    // First pass: create copies of all nodes
    curr := head
    for curr != nil {
        oldToNew[curr] = &Node{Val: curr.Val}
        curr = curr.Next
    }
    
    // Second pass: set next and random pointers
    curr = head
    for curr != nil {
        newNode := oldToNew[curr]
        if curr.Next != nil {
            newNode.Next = oldToNew[curr.Next]
        }
        if curr.Random != nil {
            newNode.Random = oldToNew[curr.Random]
        }
        curr = curr.Next
    }
    
    return oldToNew[head]
}