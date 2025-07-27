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
	mutex    sync.RWMutex
}

// введена, чтобы можно было адерсоваться на ключи мапы, вместо перебора при удалении
type listElement struct {
	key   Key
	value any
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if el, ok := c.items[key]; ok {
		val, _ := el.Value.(listElement)
		val.value = value
		el.Value = val
		c.items[key] = el
		c.queue.MoveToFront(el)
		return true
	}
	if c.queue.Len() >= c.capacity {
		last := c.queue.Back().Value.(listElement)
		c.queue.Remove(c.queue.Back())
		delete(c.items, Key(last.key))
	}
	temp := c.queue.PushFront(listElement{
		key:   key,
		value: value,
	})
	c.items[key] = temp
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	val, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(val)
		return val.Value.(listElement).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	clear(c.items)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
