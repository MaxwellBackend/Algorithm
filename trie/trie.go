package trie

import (
	"strings"
	"fmt"
)

// 屏蔽字树
type SensitiveTree struct {
	root *SensitiveNode // 根节点
}

// 屏蔽字树节点
type SensitiveNode struct {
	value    rune
	children map[rune]*SensitiveNode // 子节点
}

// 创建树
func NewSensitiveTree() *SensitiveTree {
	return &SensitiveTree{
		root: newSensitiveNode(' '),
	}
}

// 创建节点
func newSensitiveNode(value rune) *SensitiveNode {
	return &SensitiveNode{
		value:    value,
		children: make(map[rune]*SensitiveNode),
	}
}

// 插入节点
func (st *SensitiveTree) Insert(word string) {
	if strings.TrimSpace(word) == "" {
		return
	}

	node := st.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = newSensitiveNode(char)
		}
		node = node.children[char]
	}
}

// 获取根节点
func (st *SensitiveTree) GetRoot() *SensitiveNode {
	return st.root
}

// 判断节点内容是否为指定内容
func (sn *SensitiveNode) Contains(char rune) bool {
	_, found := sn.children[char]
	return found
}

// 获取指定的子节点
func (sn *SensitiveNode) GetChildNode(char rune) *SensitiveNode {
	return sn.children[char]
}

// 是否存在子节点
func (sn *SensitiveNode) HasChild() bool {
	return len(sn.children) > 0
}

// Echo
func (st *SensitiveTree) Echo() {
	st.echo(st.root)
}

func (st *SensitiveTree) echo(node *SensitiveNode) {
	for _, n := range node.children {
		fmt.Printf("%v", string(n.value))
		if len(n.children) == 0 {
			fmt.Println()
		} else {
			fmt.Print(" -> ")
		}
		st.echo(n)
	}
}
