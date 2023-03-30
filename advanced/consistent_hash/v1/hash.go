package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

type HashFunc func([]byte) uint32

type ConsistentHash struct {
	hashFunc HashFunc
	hashRing []uint32
	nodes    map[uint32]string // 节点对应其哈希值
	replicas int               // 虚拟节点数
}

// 创建一致性哈希对象
func NewConsistentHash(replicas int, hashFunc HashFunc) *ConsistentHash {
	if hashFunc == nil {
		hashFunc = Hash
	}
	return &ConsistentHash{
		hashFunc: hashFunc,
		nodes:    make(map[uint32]string),
		replicas: replicas,
	}
}

// 添加节点
func (c *ConsistentHash) AddNode(node string) {
	for i := 0; i < c.replicas; i++ {
		hash := c.hashFunc([]byte(fmt.Sprintf("%s-%d", node, i)))
		c.hashRing = append(c.hashRing, hash)
		c.nodes[hash] = node
	}
	sort.Slice(c.hashRing, func(i, j int) bool { return c.hashRing[i] < c.hashRing[j] }) // 对哈希环进行排序
}

// 移除节点
func (c *ConsistentHash) RemoveNode(node string) {
	for i := 0; i < c.replicas; i++ {
		hash := c.hashFunc([]byte(fmt.Sprintf("%s-%d", node, i)))
		delete(c.nodes, hash)
		for j, val := range c.hashRing {
			if val == hash {
				c.hashRing = append(c.hashRing[:j], c.hashRing[j+1:]...)
				break
			}
		}
	}
}

// 获取下一个节点
func (c *ConsistentHash) GetNode(key string) string {
	if len(c.hashRing) == 0 {
		return ""
	}
	hash := c.hashFunc([]byte(key))
	idx := sort.Search(len(c.hashRing), func(i int) bool { return c.hashRing[i] >= hash })
	if idx == len(c.hashRing) {
		idx = 0
	}
	return c.nodes[c.hashRing[idx]]
}

// 使用FNV算法计算哈希值
func Hash(data []byte) uint32 {
	h := fnv.New32a()
	h.Write(data)
	return h.Sum32()
}

func main() {
	ch := NewConsistentHash(3, nil)
	ch.AddNode("server1")
	ch.AddNode("server2")
	ch.AddNode("server3")
	fmt.Println(ch.GetNode("key1"))
	fmt.Println(ch.GetNode("key2"))
	ch.RemoveNode("server2")
	fmt.Println(ch.GetNode("key1"))
	fmt.Println(ch.GetNode("key2"))
}
