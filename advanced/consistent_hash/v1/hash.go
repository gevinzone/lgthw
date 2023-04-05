package v1

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type ConsistentHash struct {
	nodes       map[uint32]string // 节点哈希值与节点名称的映射
	circle      []uint32          // 哈希环
	virtualNode int               // 虚拟节点倍数
	mutex       sync.RWMutex      // 读写锁
}

// NewConsistentHash 创建 ConsistentHash 实例
func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		nodes:       make(map[uint32]string),
		virtualNode: 20,
	}
}

// AddNode 添加节点
func (c *ConsistentHash) AddNode(node string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := 0; i < c.virtualNode; i++ {
		hash := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%s%d", node, i)))
		c.circle = append(c.circle, hash)
		c.nodes[hash] = node
	}
	sort.Slice(c.circle, func(i, j int) bool { return c.circle[i] < c.circle[j] })
}

// RemoveNode 删除节点
func (c *ConsistentHash) RemoveNode(node string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := 0; i < c.virtualNode; i++ {
		hash := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%s%d", node, i)))
		delete(c.nodes, hash)
		for j := 0; j < len(c.circle); j++ {
			if c.circle[j] == hash {
				c.circle = append(c.circle[:j], c.circle[j+1:]...)
				break
			}
		}
	}
}

// GetNode 获取节点
func (c *ConsistentHash) GetNode(key string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if len(c.circle) == 0 {
		return ""
	}

	hash := crc32.ChecksumIEEE([]byte(key))
	idx := sort.Search(len(c.circle), func(i int) bool { return c.circle[i] >= hash })
	if idx == len(c.circle) {
		idx = 0
	}
	return c.nodes[c.circle[idx]]
}
