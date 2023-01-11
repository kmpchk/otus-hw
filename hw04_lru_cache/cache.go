package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newCacheElem := cacheItem{
		key:   key,
		value: value,
	}

	if curItem, ok := l.items[key]; ok {
		curItem.Value = newCacheElem
		l.queue.MoveToFront(curItem)
		return true
	}

	listElemAdded := l.queue.PushFront(newCacheElem)
	l.items[key] = listElemAdded
	if l.queue.Len() > l.capacity {
		lastQueueElemAdded := l.queue.Back().Value.(cacheItem)
		delete(l.items, lastQueueElemAdded.key)
		l.queue.Remove(l.queue.Back())
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if curItem, ok := l.items[key]; ok {
		l.queue.MoveToFront(curItem)
		elemVal := curItem.Value.(cacheItem)
		return elemVal.value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for k, v := range l.items {
		cacheElem := l.items[k].Value.(cacheItem)
		delete(l.items, cacheElem.key)
		l.queue.Remove(v)
	}
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
