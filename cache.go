package cache

type Cache[K comparable, V any] struct {
	capacity int
	list List[v]
	table    map[K]*list.Node[KV[K, V]]
	evictCb  func(key K, val V)
}

type KV[K comparable, V any] struct {
	Key K
	Val V
}

func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		lru:      list.List[KV[K, V]]{},
		table:    make(map[K]*list.Node[KV[K, V]]),
	}
}