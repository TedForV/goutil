package strings

import "log"

// Set default root value
const (
	ROOT_RUNE rune = 9999
)

// TrieNode is  trie node model
type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
	Value    rune
}

// Trie tree model,started with root node
type Trie struct {
	Root *TrieNode
}

// NewTrieNode is a new func for a trie node
func NewTrieNode(value rune) *TrieNode {
	node := new(TrieNode)
	node.Children = make(map[rune]*TrieNode)
	node.End = false
	node.Value = value
	return node
}

// NewTrie is a new func for a trie tree
func NewTrie() Trie {
	var r Trie
	r.Root = NewTrieNode(ROOT_RUNE)
	return r
}

// InsertKey is a  func that insert key into the trie tree
func (t *Trie) InsertKey(key string) {
	if len(key) == 0 {
		return
	}
	node := t.Root
	keyword := []rune(key)
	for i := 0; i < len(keyword); i++ {
		if _, existed := node.Children[keyword[i]]; !existed {
			node.Children[keyword[i]] = NewTrieNode(keyword[i])
		}
		node = node.Children[keyword[i]]
	}
	node.End = true
}

// IsExisted is a func that return the existed key
func (t *Trie) IsExisted(content string) (bool, string) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
			log.Printf("input: %s", content)
		}
	}()
	if len(content) == 0 {
		return false, ""
	}
	node := t.Root
	words := []rune(content)
	for i := 0; i < len(words); i++ {
		if existed, keys := node.isMatched(words[i:]); existed {
			return true, string(reverse(keys))
		}
	}
	return false, ""

}

//func getKeyword(tn *TrieNode) string {
//	keys := make([]rune, 10)
//	currTN := tn
//	keys = append(keys, currTN.Value)
//	for {
//		keys = append(keys, currTN.Value)
//		if tn.Parent != nil {
//			currTN = tn.Parent
//		} else {
//			break
//		}
//	}
//	keys = reverse(keys)
//	return string(keys)
//}

func (tn *TrieNode) isMatched(words []rune) (bool, []rune) {
	if child, ok := tn.Children[words[0]]; ok {
		if child.End {
			return true, []rune{words[0]}
		} else {
			if len(words) == 1 {
				return false, nil
			} else {
				result, keys := child.isMatched(words[1:])
				if result {
					return true, append(keys, words[0])
				} else {
					return false, nil
				}
			}
		}
	} else {
		return false, nil
	}

	//if  tn.End {
	//	return true, []rune{tn.Value}
	//}
	//if _, existed := tn.Children[words[0]]; existed {
	//	if len(words) == 1 {
	//		if tn.Children[words[0]].End == true {
	//			return true, words[]
	//		} else {
	//			return false, nil
	//		}
	//	}
	//
	//	result, keys := tn.Children[words[0]].isMatched(words[1:])
	//	if result {
	//		return result, append(keys, words[0])
	//
	//	} else {
	//		return false, nil
	//	}
	//	//if result {
	//	//	if tn.End == true {
	//	//		return result, keys
	//	//	} else {
	//	//		return result, append(keys, tn.Value)
	//	//	}
	//	//
	//	//} else {
	//	//	return false, nil
	//	//}
	//} else {
	//	return false, nil
	//}
}
