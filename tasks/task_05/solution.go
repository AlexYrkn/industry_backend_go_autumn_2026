package main

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{capacity: capacity, items: make(map[K]V, capacity)}
}

func (c *Cache[K, V]) Get(k K) (v V, ok bool) {
	v, ok = c.items[k]
	return
}

func (c *Cache[K, V]) Set(k K, v V) bool {
	_, exists := c.items[k]
	if len(c.items) >= c.capacity && !exists {
		return false
	}
	c.items[k] = v
	return true
}
