package cache

type Cache[K comparable, V any] struct {
	capacity int
	list List[KV[K,V]]
	table    map[K]*Node[KV[K, V]]
	evictCb  func(key K, val V)
}

type KV[K comparable, V any] struct {
	Key K
	Val V
}

// New provides initialization of the cache
func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		list:      List[KV[K, V]]{},
		table:    make(map[K]*Node[KV[K, V]]),
	}
}

// Set data to the cache
func (c *Cache[K, V]) Set(k K, v V) error {
	if d, ok := c.table[k]; ok {
		d.Value.Val = v
		c.list.Remove(d)
		c.list.PushTopNode(d)
		return nil
	}

	if len(c.table) == c.capacity {
		c.evict()
	}

	n := &Node[KV[K, V]]{
		Value: KV[K, V]{
			Key: k,
			Val: v,
		},
	}
	c.list.PushTopNode(n)
	c.table[k] = n
	return nil
}

// Get provides getting of the data from cache
func (c *Cache[K, V]) Get(k K) (V, bool) {
	if n, ok := c.table[k]; ok {
		c.list.Remove(n)
		c.list.PushTopNode(n)
		return n.Value.Val, true
	}
	var v V
	return v, false
}

// Size returns size of the cache
func (c *Cache[K, V]) Size() int {
	return len(c.table)
}

// Capacity returns capacity of the cache
func (c *Cache[K, V]) Capacity() int {
	return c.capacity
}

// Remove provides removing from cache
func (c *Cache[K, V]) Remove(k K) {
	if n, ok := c.table[k]; ok {
		c.list.Remove(n)
		delete(c.table, k)
	}
}

func (c *Cache[K, V]) evict() {
	entry := c.list.Back.Value
	if c.evictCb != nil {
		c.evictCb(entry.Key, entry.Val)
	}
	c.list.Remove(c.list.Back)
	delete(c.table, entry.Key)
}


