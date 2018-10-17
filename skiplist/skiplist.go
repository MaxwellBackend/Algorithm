package skiplist

import (
	"fmt"
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 32

// 跳表
type SkipList struct {
	header   *skipListNode
	tail     *skipListNode
	node_map map[interface{}]*skipListNode
	level    int
	Less     func(interface{}, interface{}) bool
}

type skipListLevel struct {
	forward *skipListNode
	span    int
}

type skipListNode struct {
	key      interface{}
	data     interface{}
	backward *skipListNode
	levels   []skipListLevel
}

/**************************public**************************/

func NewSkipList() *SkipList {
	sl := &SkipList{
		level:    1,
		node_map: make(map[interface{}]*skipListNode),
		header:   newSkipListNode(SKIPLIST_MAXLEVEL, 0, 0),
	}

	sl.header.backward = nil
	sl.tail = nil

	return sl
}

// 获取Level
func (s *SkipList) Level() int {
	return s.level
}

// 获取排行榜数据数量
func (s *SkipList) Length() int {
	return len(s.node_map)
}

// 根据Key设置数据
func (s *SkipList) Set(key interface{}, data interface{}) error {
	if e := s.Delete(key); e != nil {
		return e
	}

	var update [SKIPLIST_MAXLEVEL]*skipListNode
	var rank [SKIPLIST_MAXLEVEL]int
	var node *skipListNode

	node = s.header

	for i := s.level - 1; i >= 0; i-- {
		if i == s.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for node.levels[i].forward != nil && s.Less(node.levels[i].forward.data, data) {
			rank[i] += node.levels[i].span
			node = node.levels[i].forward
		}
		update[i] = node
	}

	var level = randomLevel()
	if level > s.level {
		for i := s.level; i < level; i++ {
			rank[i] = 0
			update[i] = s.header
			update[i].levels[i].span = s.Length()
		}
		s.level = level
	}

	node = newSkipListNode(level, key, data)
	s.node_map[key] = node

	for i := 0; i < level; i++ {
		node.levels[i].forward = update[i].levels[i].forward
		update[i].levels[i].forward = node

		node.levels[i].span = update[i].levels[i].span - (rank[0] - rank[i])
		update[i].levels[i].span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < s.level; i++ {
		update[i].levels[i].span++
	}

	if update[0] == s.header {
		node.backward = nil
	} else {
		node.backward = update[0]
	}

	if node.levels[0].forward != nil {
		node.levels[0].forward.backward = node
	} else {
		s.tail = node
	}

	return nil
}

// 根据Key获取数据
func (s *SkipList) Get(key interface{}) interface{} {
	if node := s.getNode(key); node != nil {
		return node.data
	}

	return nil
}

// 根据Key删除数据
func (s *SkipList) Delete(key interface{}) error {
	if node := s.getNode(key); node != nil {
		return s.deleteByNode(node)
	}

	return nil
}

// 根据Key判断是否存在
func (s *SkipList) Exist(key interface{}) bool {
	return s.Get(key) != nil
}

// 根据Key获取排名
func (s *SkipList) GetRank(key interface{}) int {
	if node := s.getNode(key); node != nil {
		return s.GetRankByData(node.data)
	}

	return 0
}

// 根据数据获取排名
func (s *SkipList) GetRankByData(data interface{}) int {
	var node *skipListNode
	var rank int = 1

	node = s.header
	for i := s.level - 1; i >= 0; i-- {
		for node.levels[i].forward != nil && s.Less(node.levels[i].forward.data, data) {
			rank += node.levels[i].span
			node = node.levels[i].forward
		}
	}

	node = node.levels[0].forward
	if node != nil && data == node.data {
		return rank
	}

	return 0
}

// 根据排名获取数据
func (s *SkipList) GetDataByRank(rank int) interface{} {
	if rank == 0 || rank > s.Length() {
		return nil
	}

	var node *skipListNode
	var traversed int = 0

	node = s.header
	for i := s.level - 1; i >= 0; i-- {
		for node.levels[i].forward != nil && (traversed+node.levels[i].span) <= rank {
			traversed += node.levels[i].span
			node = node.levels[i].forward
		}

		if traversed == rank {
			return node.data
		}
	}

	return nil
}

// 获取排名前多少名数据
func (s *SkipList) Top(number int) []interface{} {
	var top_list = make([]interface{}, number)
	var index int

	node := s.header.levels[0].forward
	for node != nil && index < number {
		top_list[index] = node.data
		index++
		node = node.levels[0].forward
	}

	return top_list
}

// 获取所有数据
func (s *SkipList) All() []interface{} {
	var top_list = []interface{}{}
	var index int

	node := s.header.levels[0].forward
	for node != nil {
		top_list = append(top_list, node.data)
		index++
		node = node.levels[0].forward
	}

	return top_list
}

// 获取排名范围内的数据
func (s *SkipList) Range(min, max int) []interface{} {
	var top_list = make([]interface{}, max-min)
	var index int

	node := s.header.levels[0].forward
	for node != nil {
		if index >= max {
			break
		}

		if index >= min {
			top_list[index-min] = node.data
		}

		node = node.levels[0].forward
		index++
	}

	return top_list
}

// 获取排名后多少名数据
func (s *SkipList) Bottom(number int) []interface{} {
	var bottom_list = make([]interface{}, number)
	var index int

	node := s.tail
	for node != nil && index < number {
		bottom_list[index] = node.data
		index++
		node = node.backward
	}

	return bottom_list
}

// 清空排行榜
func (s *SkipList) Clear() {
	s.header = newSkipListNode(SKIPLIST_MAXLEVEL, 0, 0)
	s.tail = nil
	s.node_map = make(map[interface{}]*skipListNode)
	s.level = 1
}

/**************************private**************************/

func newSkipListNode(level int, key interface{}, data interface{}) *skipListNode {
	sn := &skipListNode{
		key:    key,
		data:   data,
		levels: make([]skipListLevel, level),
	}
	return sn
}

// 随机节点高度
func randomLevel() int {
	var level int = 1
	for {
		if rand.Intn(2) == 1 {
			level++

			if level >= SKIPLIST_MAXLEVEL {
				break
			}

		} else {
			break
		}
	}

	return level
}

// 根据数据删除节点
func (s *SkipList) deleteByData(data interface{}) error {
	var update = make([]*skipListNode, SKIPLIST_MAXLEVEL)
	var node *skipListNode

	node = s.header
	for i := s.level - 1; i >= 0; i-- {
		for node.levels[i].forward != nil && s.Less(node.levels[i].forward.data, data) {
			node = node.levels[i].forward
		}
		update[i] = node
	}

	node = node.levels[0].forward
	if data == node.data {
		s.deleteNode(node, update)
		return nil
	}

	return fmt.Errorf("deleteByData not found data : %v", data)
}

// 根据节点删除
func (s *SkipList) deleteByNode(node *skipListNode) error {
	if node == nil {
		return nil
	}

	return s.deleteByData(node.data)
}

// 删除节点
func (s *SkipList) deleteNode(x *skipListNode, update []*skipListNode) {
	for i := 0; i < s.level; i++ {
		if update[i].levels[i].forward == x {
			update[i].levels[i].span += x.levels[i].span - 1
			update[i].levels[i].forward = x.levels[i].forward
		} else {
			update[i].levels[i].span -= 1
		}
	}

	if x.levels[0].forward != nil {
		x.levels[0].forward.backward = x.backward
	} else {
		s.tail = x.backward
	}

	for s.level > 1 && s.header.levels[s.level-1].forward == nil {
		s.level--
	}

	delete(s.node_map, x.key)
}

// 获取节点
func (s *SkipList) getNode(key interface{}) *skipListNode {
	return s.node_map[key]
}
