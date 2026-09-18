type Node struct {
	key, val int
	prev, next *Node
}
type LRUCache struct {
    cache map[int]*Node
	cap int
	head,tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
    return LRUCache{
		cache: make(map[int]*Node, capacity),
		cap: capacity,
		head: head, // LRU
		tail: tail, // MRU
	}
}

func (this *LRUCache) swapNeighbor(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) pushToTail(node *Node) {
	node.next, node.prev = this.tail, this.tail.prev
	this.tail.prev.next = node
	this.tail.prev = node
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {		
		this.swapNeighbor(node)
		this.pushToTail(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.val = value
		this.swapNeighbor(node)
		this.pushToTail(node)
		return
	}

	if len(this.cache) == this.cap {
		lru := this.head.next
		this.swapNeighbor(lru)
		delete(this.cache, lru.key)
	}
	newNode := &Node{val: value, key: key}
	this.pushToTail(newNode)
	this.cache[key] = newNode
}
