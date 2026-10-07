package p0146lrucache

type node struct {
	key, val   int
	prev, next *node
}

type LRUCache struct {
	cap        int
	cache      map[int]*node
	head, tail *node
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*node),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) pushFront(n *node) {
	n.prev = this.head
	n.next = this.head.next
	this.head.next.prev = n
	this.head.next = n

}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(n)
	this.pushFront(n)
	return n.val

}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.cache[key]; ok {
		n.val = value
		this.remove(n)
		this.pushFront(n)
		return
	}

	n := &node{key: key, val: value}
	this.cache[key] = n
	this.pushFront(n)

	if len(this.cache) > this.cap {
		lru := this.tail.prev // node paling lama
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
