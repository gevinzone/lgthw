package hash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	ch := NewConsistentHash()

	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	//打印哈希环上所有虚拟节点和真实节点的对应关系
	for k, v := range ch.circle {
		t.Logf("虚拟节点哈希值：%d，节点：%d\n", k, v)
	}

	//打印获取数据存储节点的名称
	for i := 0; i < 10; i++ {
		t.Logf("数据 %d 存储节点：%s\n", i, ch.GetNode(strconv.Itoa(i)))
	}

	ch.RemoveNode("node2")

	//打印删除一个节点后，哈希环上所有虚拟节点和真实节点的对应关系
	fmt.Println("删掉一个节点后的虚拟节点和真实节点的对应关系：")
	for k, v := range ch.circle {
		t.Logf("虚拟节点哈希值：%d，节点：%d\n", k, v)
	}

	//打印获取数据存储节点的名称
	for i := 0; i < 10; i++ {
		t.Logf("数据 %d 存储节点：%s\n", i, ch.GetNode(strconv.Itoa(i)))
	}
}
