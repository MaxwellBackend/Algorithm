package skiplist


import (
	"math/rand"
	"fmt"
)

const SKIPLIST_MAXLEVEL = 8

type Node struct {
	Forward  []Node
	Value interface{}
}

func NewNode(v interface{},level int) *Node {
	return &Node{
		Value: v,
		Forward: make([]Node, level),
	}
}

type SkipList struct {
	Header *Node
	Level  int
}

func NewSkipList() *SkipList {
	return &SkipList{Level: 1, Header: NewNode(0,SKIPLIST_MAXLEVEL)}
}

func (skipList *SkipList) Insert(key int) {

	update := make(map[int]*Node)
	node := skipList.Header

	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			if node.Forward[i].Value != nil && node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
			}else{
				break;
			}
		}
		update[i] = node
	}

	level := skipList.RandomLevel()
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}

	newNode := NewNode(key,level)

	for i := 0; i < level; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = *newNode
	}

}

func (skipList *SkipList) Remove(key int) {

	update := make(map[int]*Node)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			if node.Forward[i].Value == nil {
				break
			}
			if node.Forward[i].Value.(int) == key {
				update[i] = node
				break
			}
			if  node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			}else{
				break
			}
		}
	}

	for i,v := range update {
		if v == skipList.Header {
			skipList.Level --
		}
		v.Forward[i] = v.Forward[i].Forward[i]
	}
}

func (skipList *SkipList) Search(key int) *Node {

	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				return &node.Forward[i]
			}

			if  node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else {
				break
			}
		}
	}
	return nil
}

func (skipList *SkipList) RandomLevel() int {
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

func (skipList *SkipList) PrintSkipList() {

	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {

		fmt.Println("level:", i)
		node := skipList.Header.Forward[i]
		for {
			if node.Value != nil {
				fmt.Printf("%d ", node.Value.(int))
				node = node.Forward[i]
			}else{
				break
			}
		}
		fmt.Println("\n--------------------------------------------------------")
	}

	fmt.Println("Current MaxLevel:", skipList.Level)
}
