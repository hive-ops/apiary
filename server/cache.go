package server

import (
	"github.com/hive-ops/apiary/utils"
	"sync"
)

type CacheConfig struct {
	TTL            int `json:"ttl"`
	MaxObjectCount int `json:"max_object_count"`
}

type Cache struct {
	config CacheConfig
	store  *utils.HashMap
	keys   *utils.DoublyLinkedList
	sync.RWMutex
}

func NewCache() *Cache {

	config := CacheConfig{
		TTL:            60,
		MaxObjectCount: 1000,
	}

	return &Cache{
		config: config,
		store:  utils.NewHashMap(),
		keys:   utils.NewDoublyLinkedList(config.MaxObjectCount),
	}
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.RLock()
	node, ok := c.store.Get(key)
	if !ok {
		c.RUnlock()
		return nil, nil
		//return "", errors.New("not found")
	}
	c.keys.MoveToFront(node)
	c.RUnlock()

	return node.Value, nil
}

func (c *Cache) Set(key string, value []byte) {
	c.Lock()
	if node, ok := c.store.Get(key); ok {
		node.Value = value
		c.keys.MoveToFront(node)
		c.Unlock()
		return
	}
	if c.keys.Size >= c.config.MaxObjectCount {
		nodeToDelete := c.keys.RemoveLast()
		c.store.Delete(nodeToDelete.Key)
		c.keys.Remove(nodeToDelete)
	}
	node := c.keys.PushFront(key, value)
	c.store.Set(key, node)

	c.Unlock()
}

func (c *Cache) Delete(key string) {
	c.Lock()
	if nodeToDelete, ok := c.store.Get(key); ok {
		c.keys.Remove(nodeToDelete)
		c.store.Delete(key)
	}
	c.Unlock()
}

func (c *Cache) Clear() {
	c.Lock()
	c.store.Clear()
	c.keys.Clear()
	c.Unlock()
}
