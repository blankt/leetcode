package leetcode

type DoubleLinkedListNode struct {
	key  int
	val  int
	prev *DoubleLinkedListNode
	next *DoubleLinkedListNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*DoubleLinkedListNode
	head     *DoubleLinkedListNode
	tail     *DoubleLinkedListNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
	}
	cache.cache = make(map[int]*DoubleLinkedListNode)
	cache.head = &DoubleLinkedListNode{}
	cache.tail = &DoubleLinkedListNode{}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	if this.head.next == node {
		return node.val
	}
	this.remove(node)
	this.add(node.key, node.val)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		this.remove(node)
		this.add(key, value)
	} else {
		if len(this.cache) >= this.capacity {
			this.remove(this.tail.prev)
		}
		this.add(key, value)
	}
}

func (this *LRUCache) remove(node *DoubleLinkedListNode) {
	if node == this.head || node == this.tail {
		return
	}
	node.next.prev = node.prev
	node.prev.next = node.next
	delete(this.cache, node.key)
}

func (this *LRUCache) add(key int, value int) {
	node := &DoubleLinkedListNode{val: value, key: key}
	this.cache[key] = node
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
