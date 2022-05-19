package cache

import (
	"errors"
	"fmt"
	"sync"
)

type cache struct {
	storage map[string]string
	m       sync.RWMutex
}

func New() *cache {
	return &cache{storage: make(map[string]string, 1024), m: sync.RWMutex{}}
}

func (c *cache) Set(key, value string) {
	c.m.Lock()
	defer c.m.Unlock()
	if _, ok := c.storage[key]; !ok {
		c.storage[key] = value
	}
}

func (c *cache) Get(key string) (data string, err error) {
	c.m.RLock()
	if v, ok := c.storage[key]; ok {
		return v, nil
	}
	return data, errors.New("not found")
}

func (c *cache) Print() {
	fmt.Println(c.storage)
}
