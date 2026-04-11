package daily

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	Cache    map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{Key: 0, Value: 0}
	tail := &Node{Key: 0, Value: 0}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node),
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) remove(node *Node) {
	prevNode := node.Prev
	nextNode := node.Next

	prevNode.Next = nextNode
	nextNode.Prev = prevNode
}

func (this *LRUCache) add(node *Node) {
	headFirst := this.Head.Next
	node.Next = headFirst
	node.Prev = this.Head
	this.Head.Next = node
	headFirst.Prev = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.remove(node)
		this.add(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		this.remove(node)
	}
	newNode := &Node{Key: key, Value: value}
	this.Cache[key] = newNode
	this.add(newNode)

	if len(this.Cache) > this.Capacity {
		lruNode := this.Tail.Prev
		this.remove(lruNode)
		delete(this.Cache, lruNode.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
