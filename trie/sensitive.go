package trie

import (
	"strings"
	"os"
	"bufio"
	"io"
)

var tree *SensitiveTree

// 构建屏蔽字树
func CreateSensitiveTree() error {
	f, err := os.Open("../sensitive.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	tree = NewSensitiveTree()

	reader := bufio.NewReader(f)
	for {
		lstr, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			} else {
				break
			}
		}

		strs := strings.Split(strings.ToLower(string(lstr)), "|")
		if len(strs) <= 0 {
			continue
		}

		str := strs[0]
		if len(str) <= 0 {
			continue
		}
		firstWord := []rune(str)[0]
		for i := 1; i < len(strs); i++ {
			word := string(firstWord) + strs[i]
			tree.Insert(word)
		}
	}

	return nil
}

// 转换
func SensitiveTransform(word string) string {

	t := []rune(strings.ToLower(word)) // 转换为unicode数组
	p1 := tree.GetRoot()               // p1指向树的根节点
	p2 := t[0]                         // p2是过滤内容的第一个字符
	var p2Index, p3Index int           // p2的位置,p3的位置,默认是过滤内容的第一个字符的位置
	newWord := word                    // 过滤后的字符串

	// p3指向最后一个字符，则结束
	for ; p3Index < len(t); {

		// 遍历从p3开始到结尾的字符串
		for i := p3Index; i < len(t); i++ {
			char := t[i]

			if !p1.Contains(p2) {

				// 如果p2不是p1子节点，说明以p3开始的内容不需要屏蔽，p3指向下一个位置，p2指向p3，并将p1重置到root节点，终止循环
				p2Index = p3Index + 1
				p3Index = p3Index + 1

				if p2Index < len(t) {
					p2 = t[p2Index]
				}
				p1 = tree.GetRoot()

				break
			} else {

				// 如果p2是p1的子节点，p1节点移动到p2内容对应的节点上，p2移动到下个位置

				// 获取p1的子节点
				p1 = p1.GetChildNode(char)

				// 移动P2到下个位置
				p2Index = p2Index + 1
				if p2Index < len(t) {
					p2 = t[p2Index]
				}

				if !p1.HasChild() {

					// p1没有子节点，树的某条路径遍历完了，说明字符串中含有敏感词
					p1 = tree.GetRoot()

					// 将敏感词替换成*
					temp := []rune(newWord)
					newWord = string(temp[:p3Index]) + strings.Repeat("*", p2Index-p3Index) + string(temp[p2Index:])

					// p3移动到p2的位置
					p3Index = p2Index
				}
			}
		}
	}

	return newWord
}

// 检测
func SensitiveCheck(word string) bool {
	t := []rune(strings.ToLower(word)) // 转换为unicode数组
	p1 := tree.GetRoot()               // p1指向树的根节点
	p2 := t[0]                         // p2是过滤内容的第一个字符
	var p2Index, p3Index int           // p2的位置,p3的位置,默认是过滤内容的第一个字符的位置

	// p3指向最后一个字符，则结束
	for ; p3Index < len(t); {

		// 遍历从p3开始到结尾的字符串
		for i := p3Index; i < len(t); i++ {
			char := t[i]

			if !p1.Contains(p2) {

				// 如果p2不是p1子节点，说明以p3开始的内容不需要屏蔽，p3指向下一个位置，p2指向p3，并将p1重置到root节点，终止循环
				p2Index = p3Index + 1
				p3Index = p3Index + 1

				if p2Index < len(t) {
					p2 = t[p2Index]
				}
				p1 = tree.GetRoot()

				break
			} else {

				// 如果p2是p1的子节点，p1节点移动到p2内容对应的节点上，p2移动到下个位置

				// 获取p1的子节点
				p1 = p1.GetChildNode(char)

				// 移动P2到下个位置
				p2Index = p2Index + 1
				if p2Index < len(t) {
					p2 = t[p2Index]
				}

				if !p1.HasChild() {

					// p1没有子节点，树的某条路径遍历完了，说明字符串中含有敏感词
					p1 = tree.GetRoot()

					return true

					// p3移动到p2的位置
					p3Index = p2Index
				}
			}
		}
	}

	return false
}
