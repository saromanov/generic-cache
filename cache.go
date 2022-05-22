package cache

type Cache[K comparable, V any] struct {
	capacity int
	table    map[K]*list.Node[KV[K, V]]
	evictCb  func(key K, val V)
}