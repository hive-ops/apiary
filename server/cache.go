package server

import "errors"

type Cache struct {
	entries map[string]string
}

func NewCache() Cache {
	return Cache{
		entries: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, error) {
	value, ok := c.entries[key]
	if !ok {
		return "", errors.New("not found")
	}
	return value, nil
}

func (c *Cache) Set(key, value string) {
	c.entries[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.entries, key)
}
