package daily

type Node struct {
	key, val   int
	freq       int
	prev, next *Node
}

// ─────────────────────────────────────────
// DLinkedList
// ─────────────────────────────────────────

type DLinkedList struct {
	head, tail *Node
	size       int
}

func newDLinkedList() *DLinkedList {
	dll := &DLinkedList{
		head: &Node{},
		tail: &Node{},
	}
	dll.head.next = dll.tail
	dll.tail.prev = dll.head
	return dll
}

func (dll *DLinkedList) addFront(node *Node) {
	nxt := dll.head.next
	dll.head.next = node
	node.prev = dll.head
	node.next = nxt
	nxt.prev = node
	dll.size++
}

func (dll *DLinkedList) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	dll.size--
}

func (dll *DLinkedList) removeLast() *Node {
	if dll.size == 0 {
		return nil
	}
	last := dll.tail.prev
	dll.remove(last)
	return last
}

// ─────────────────────────────────────────
// LFUCache
// ─────────────────────────────────────────

type LFUCache struct {
	cap       int
	minFreq   int
	keyToNode map[int]*Node
	buckets   map[int]*DLinkedList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		keyToNode: make(map[int]*Node),
		buckets:   make(map[int]*DLinkedList),
	}
}

// getBucket replaces Python's defaultdict — creates the bucket on first access
func (lfu *LFUCache) getBucket(freq int) *DLinkedList {
	if _, ok := lfu.buckets[freq]; !ok {
		lfu.buckets[freq] = newDLinkedList()
	}
	return lfu.buckets[freq]
}

func (lfu *LFUCache) promote(node *Node) {
	lfu.getBucket(node.freq).remove(node)
	if lfu.buckets[node.freq].size == 0 && node.freq == lfu.minFreq {
		lfu.minFreq++
	}
	node.freq++
	lfu.getBucket(node.freq).addFront(node)
}

func (lfu *LFUCache) Get(key int) int {
	node, ok := lfu.keyToNode[key]
	if !ok {
		return -1
	}
	lfu.promote(node)
	return node.val
}

func (lfu *LFUCache) Put(key, value int) {
	if lfu.cap == 0 {
		return
	}
	if node, ok := lfu.keyToNode[key]; ok {
		node.val = value
		lfu.promote(node)
	} else {
		if len(lfu.keyToNode) >= lfu.cap {
			evicted := lfu.getBucket(lfu.minFreq).removeLast()
			delete(lfu.keyToNode, evicted.key)
		}
		newNode := &Node{key: key, val: value, freq: 1}
		lfu.keyToNode[key] = newNode
		lfu.getBucket(1).addFront(newNode)
		lfu.minFreq = 1
	}
}
