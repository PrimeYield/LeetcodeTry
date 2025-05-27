// 208. Implement Trie (Prefix Tree)
package trie

// type Trie struct {
// 	root *TrieNode
// }

// func NewTrie() *Trie {
// 	return &Trie{
// 		root: &TrieNode{},
// 	}
// }

// type TrieNode struct {
// 	Children  [26]*TrieNode
// 	PassCount int
// 	End       bool
// }

// func Constructor() Trie {
// 	return Trie{}
// }

// func (t *Trie) Insert(word string) {
// 	//已經存在就不需要新增
// 	if t.Search(word) {
// 		return
// 	}
// 	move := t.root
// 	for _, ch := range word {
// 		if move.Children[ch-'a'] == nil {
// 			move.Children[ch-'a'] = &TrieNode{}
// 		}
// 		move.PassCount++
// 		move = move.Children[ch-'a']
// 	}
// 	move.End = true
// }

// func (t *Trie) Search(word string) bool {
// 	node := t.search(word)
// 	if node == nil {
// 		return false
// 	}
// 	return node != nil && node.End
// }

// func (t *Trie) StartsWith(prefix string) bool {
// 	node := t.search(prefix)
// 	if node == nil {
// 		return false
// 	}
// 	return true
// }

// func (t *Trie) search(word string) *TrieNode {
// 	move := t.root
// 	for _, ch := range word {
// 		if move.Children[ch-'a'] == nil {
// 			return nil
// 		}
// 		move = move.Children[ch-'a']
// 	}
// 	return move
// }

type Trie struct {
	Children  [26]*Trie
	PassCount int
	End       bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}
	for _, ch := range word {
		if t.Children[ch-'a'] == nil {
			t.Children[ch-'a'] = &Trie{}
		}
		t.PassCount++
		t = t.Children[ch-'a']
	}
	t.End = true
}

func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.End
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.search(prefix)
	return node != nil
}

func (t *Trie) search(word string) *Trie {
	move := t
	for _, ch := range word {
		if move.Children[ch-'a'] == nil {
			return nil
		}
		move = move.Children[ch-'a']
	}
	return move
}

//runtime 32ms Memory 17.83MB
