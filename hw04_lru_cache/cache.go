package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if i, exists := c.items[key]; exists {
		i.Value.(*cacheItem).value = value
		c.queue.MoveToFront(i)
		return true
	}
	newItem := c.queue.PushFront(&cacheItem{key, value})
	c.items[key] = newItem
	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		delete(c.items, last.Value.(*cacheItem).key)
		c.queue.Remove(last)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if i, exists := c.items[key]; exists {
		c.queue.MoveToFront(i)
		return i.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
