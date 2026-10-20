package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}
type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	element, found := c.items[key]
	if !found {
		return
	}
	c.ll.MoveToFront(element)
	return element.Value.(entry[K, V]).value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	element, found := c.items[key]
	if found {
		element.Value = entry[K, V]{key, value}
		return
	}
	if c.capacity == len(c.items) {
		elementToDelete := c.ll.Back()
		c.ll.Remove(elementToDelete)
		keyToDelete := elementToDelete.Value.(entry[K, V]).key
		delete(c.items, keyToDelete)
	}
	e := c.ll.PushFront(entry[K, V]{key, value})
	c.items[key] = e
}
