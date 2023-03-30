package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type ConsistentHash struct {
	nodes       map[uint32]string // 节点哈希值与节点名称的映射
	circle      []uint32          // 哈希环
	virtualNode int               // 虚拟节点倍数
	mutex       sync.RWMutex      // 读写锁
}

// 创建 ConsistentHash 实例
func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{
		nodes:       make(map[uint32]string),
		virtualNode: 20,
	}
}

// 添加节点
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

// 删除节点
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

// 获取节点
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

//测试一致性哈希算法
func main() {
	ch := NewConsistentHash()

	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	//打印哈希环上所有虚拟节点和真实节点的对应关系
	for k, v := range ch.circle {
		fmt.Printf("虚拟节点哈希值：%d，节点：%s\n", k, v)
	}

	//打印获取数据存储节点的名称
	for i := 0; i < 10; i++ {
		fmt.Printf("数据 %d 存储节点：%s\n", i, ch.GetNode(strconv.Itoa(i)))
	}

	ch.RemoveNode("node2")

	//打印删除一个节点后，哈希环上所有虚拟节点和真实节点的对应关系
	fmt.Println("删掉一个节点后的虚拟节点和真实节点的对应关系：")
	for k, v := range ch.circle {
		fmt.Printf("虚拟节点哈希值：%d，节点：%s\n", k, v)
	}

	//打印获取数据存储节点的名称
	for i := 0; i < 10; i++ {
		fmt.Printf("数据 %d 存储节点：%s\n", i, ch.GetNode(strconv.Itoa(i)))
	}
}
