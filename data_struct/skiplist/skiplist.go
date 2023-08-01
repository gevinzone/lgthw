// Copyright 2023 igevin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package skiplist

const maxLevel = 32

// SkipList 是一个跳表，层高为levelCount, 自底向上为第0层、1层…… levelCount-1 层
// 第0层是跳表全部数据，排成一个有序链表，之上逐层创建了索引，以加速0层链表的crud，整个跳表可以视为一个索引
// 每层均是一个有序链表，上一层的节点指向下一层的相同节点，从而整个跳表的数据，前后上下均可以联通
// head 是头指针，指向整个跳表，表示为一个Node数据结构
// Node 结构中，val存储本身数据，forward 数组，存储每层上，该节点的下一个节点
// 故head的forward 数组，指向了每层的头指针
type SkipList struct {
	head       *Node
	levelCount int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:       newNode(),
		levelCount: 1,
	}
}

func (s *SkipList) Find(val int) *Node {
	p := s.head
	// 找到val 对应的索引项
	for i := s.levelCount - 1; i >= 0; i-- {
		for p.forward[i] != nil && p.forward[i].val < val {
			p = p.forward[i]
		}
	}
	// 由于上层的数据，下层一定有，故不管当前在第几层，直接在第0层确认到底有没有该数据即可
	if p.forward[0] != nil && p.forward[0].val == val {
		return p
	} else {
		return nil
	}
}

type Node struct {
	val       int
	forward   []*Node
	nMaxLevel int
}

func newNode() *Node {
	return &Node{
		val:       -1,
		forward:   make([]*Node, maxLevel),
		nMaxLevel: 0,
	}
}
