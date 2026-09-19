type LRUCache struct {
	cache    map[int]*list.Element
	list     *list.List
	capacity int
}

type Node struct {
	val int
	key int
}

// list package
// *list.List
// *list.Element - Node
// list.MoveToFront(), list.MoveToBack()
// list.Remove(), list.PushFront(), list.PushBack()
// list.Back(), list.Front()
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.list.MoveToFront(node)
		return node.Value.(Node).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.Value = Node{value, key}
		this.list.MoveToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		lru := this.list.Back()
		this.list.Remove(lru)
		delete(this.cache, lru.Value.(Node).key)
	}

	this.cache[key] = this.list.PushFront(Node{value, key})
}