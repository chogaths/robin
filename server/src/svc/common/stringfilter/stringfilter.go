package stringfilter

import (
	"unicode/utf8"
)

type StringFilter struct {
	root      *stringFilterNode // 根节点
	*ListFile                   // 文件行处理模块
}

type stringFilterNode struct {
	children map[rune]*stringFilterNode
	end      bool // 是否是单词的结束
}

//new stringfilter
func NewStringFilter() *StringFilter {
	t := new(StringFilter)
	t.root = newStringFilterNode()
	t.ListFile = new(ListFile)
	return t
}

//new stringfilterNode
func newStringFilterNode() *stringFilterNode {
	t := new(stringFilterNode)
	t.end = false
	t.children = make(map[rune]*stringFilterNode)
	return t
}

// 插入敏感词
func (self *StringFilter) Insert(txt string) {
	if len(txt) < 1 {
		return
	}

	node := self.root
	key := []rune(txt)

	// 创建trie树（又称单词查找树）
	for i := 0; i < len(key); i++ {
		if _, exists := node.children[key[i]]; !exists {
			node.children[key[i]] = newStringFilterNode()
		}
		node = node.children[key[i]]
	}

	node.end = true
}

// 替换文字、参数为要替换的文字内容
func (self *StringFilter) Replace(txt string) string {
	if len(txt) < 1 {
		return txt
	}

	node := self.root

	key := []rune(txt)
	var chars []rune = nil
	slen := len(key)

	for i := 0; i < slen; i++ {
		var match bool
		var endPos int
		if _, exists := node.children[key[i]]; exists {
			node = node.children[key[i]]
			if node.end { // 单个单词匹配
				c, _ := utf8.DecodeRuneInString("*")
				if chars == nil {
					chars = key
				}
				chars[i] = c
			}
			for j := i + 1; j < slen; j++ {
				if _, exists := node.children[key[j]]; !exists {
					break
				}

				node = node.children[key[j]]
				if !node.end {
					continue
				}

				match = true
				endPos = j

				if len(node.children) > 0 {
					continue
				}
			}

			if match {
				if chars == nil {
					chars = key
				}
				for t := i; t <= endPos; t++ { // 从敏感词开始到结束依次替换为*
					c, _ := utf8.DecodeRuneInString("*")
					chars[t] = c
				}

			}
			node = self.root
		}
	}
	if chars == nil {
		return txt
	} else {
		return string(chars)
	}
}

// 判断是否有敏感词的存在
func (self *StringFilter) Exist(txt string) bool {
	if len(txt) < 1 {
		return false
	}

	node := self.root
	key := []rune(txt)

	for i := 0; i < len(key); i++ {
		if _, exists := node.children[key[i]]; exists {
			node = node.children[key[i]]
			if node.end { // 单个单词匹配
				return true
			}

			for j := i + 1; j < len(key); j++ {
				if _, exists := node.children[key[j]]; !exists {
					break
				}

				node = node.children[key[j]]
				if !node.end {
					continue
				}
				return true
			}
			node = self.root
		}
	}
	return false
}

// 加载敏感词词库
func (self *StringFilter) LoadBlogFile(fileName string) {
	self.Load(fileName, self.Insert)
}

// 是否插入了敏感词
func (self *StringFilter) Empty() bool {
	return len(self.root.children) == 0
}
