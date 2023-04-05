package chash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type ConsistentHash struct {
	hash     HashFunc
	replicas int
	keys     []uint32
	nodesMap map[uint32]string
	mu       sync.RWMutex // 互斥锁（mutex）
}

type ConsistentHashOption func(c *ConsistentHash)

func NewConsistentHash(opts ...ConsistentHashOption) *ConsistentHash {
	c := &ConsistentHash{
		hash:     crc32.ChecksumIEEE,
		replicas: 20,
		keys:     make([]uint32, 20),
		nodesMap: make(map[uint32]string),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithHash(h HashFunc) ConsistentHashOption {
	return func(c *ConsistentHash) {
		c.hash = h
	}
}

func WithReplica(replicas int) ConsistentHashOption {
	return func(c *ConsistentHash) {
		c.replicas = replicas
	}
}

func (c *ConsistentHash) AddNode(node string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < c.replicas; i++ {
		h := c.hash([]byte(fmt.Sprintf("%s%d", node, i)))
		c.nodesMap[h] = node
		c.keys = append(c.keys, h)
	}
	sort.Slice(c.keys, func(i, j int) bool { return c.keys[i] < c.keys[j] })
}

func (c *ConsistentHash) RemoveNode(node string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < c.replicas; i++ {
		h := c.hash([]byte(fmt.Sprintf("%s%d", node, i)))
		delete(c.nodesMap, h)
		for j := 0; j < len(c.keys); j++ {
			if c.keys[j] == h {
				c.keys = append(c.keys[:j], c.keys[j+1:]...)
			}
		}
	}
}

func (c *ConsistentHash) GetNode(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	h := c.hash([]byte(key))
	idx := sort.Search(len(c.keys), func(i int) bool { return c.keys[i] >= h })
	if idx == len(c.keys) {
		idx = 0
	}
	return c.nodesMap[c.keys[idx]]
}
