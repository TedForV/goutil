package strings

import (
	"log"
	"math"
	"sync"
)

// Set default root value
const (
	RootRune rune = 9999
	MaxLen   int  = 50
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
	// MaxLevel is the max leve of trie tree.
	MaxLevel int
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
	r.MaxLevel = 0
	r.Root = NewTrieNode(RootRune)
	return r
}

// InsertKey is a  func that insert key into the trie tree
func (t *Trie) InsertKey(key string) {
	if len(key) == 0 {
		return
	}
	node := t.Root
	keyword := []rune(key)
	length := len(keyword)
	if length > t.MaxLevel {
		t.MaxLevel = length
	}
	for i := 0; i < length; i++ {
		if _, existed := node.Children[keyword[i]]; !existed {
			node.Children[keyword[i]] = NewTrieNode(keyword[i])
		}
		node = node.Children[keyword[i]]
	}
	node.End = true
}

// IsExisted is a func that return the existed key
func (t *Trie) IsExisted(content string) (bool, int, string) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
			log.Printf("input: %s", content)
		}
	}()
	if len(content) == 0 {
		return false, -1, ""
	}
	node := t.Root
	words := []rune(content)
	for i := 0; i < len(words); i++ {
		if existed, keys := node.isMatched(words[i:]); existed {
			return true, i, string(reverse(keys))
		}
	}
	return false, -1, ""

}

// FindAllExisted get all existed content and return
func (t *Trie) FindAllExisted(content string) []string {
	if len(content) == 0 {
		return []string{}
	}
	words := []rune(content)
	length := len(words)
	round := seperateContent(length)
	var wg sync.WaitGroup
	chl := make(chan string, 10)
	minIndex, maxIndex := 0, 0
	existed := make(map[string]struct{}, 10)
	for i := 0; i < round; i++ {
		if i*MaxLen > length {
			break
		}
		wg.Add(1)
		minIndex = i * MaxLen
		maxIndex = (i+1)*MaxLen + t.MaxLevel
		if maxIndex > length {
			maxIndex = length
		}
		go func(minIndex, maxIndex int) {
			t.findAllExisted(words[minIndex:maxIndex], chl)
			wg.Done()
		}(minIndex, maxIndex)
	}
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		for value := range chl {
			existed[value] = struct{}{}
		}
		wg2.Done()
	}()
	wg.Wait()
	close(chl)
	wg2.Wait()
	data := make([]string, len(existed))
	i := 0
	for k := range existed {
		data[i] = k
		i++
	}
	return data
}

func (t *Trie) findAllExisted(content []rune, chl chan<- string) {
	length := len(content)
	for i := 0; i < length; i++ {
		if existed, keys := t.Root.isMatched(content[i:]); existed {
			chl <- string(reverse(keys))
			i += len(keys) - 1
		}
	}
}

func seperateContent(contentLen int) int {
	if contentLen == 0 {
		return 0
	}
	return int(math.Ceil(float64(contentLen) / float64(MaxLen)))
}

func (tn *TrieNode) isMatched(words []rune) (bool, []rune) {
	if child, ok := tn.Children[words[0]]; ok {
		if child.End {
			return true, []rune{words[0]}
		}
		if len(words) == 1 {
			return false, nil
		}
		result, keys := child.isMatched(words[1:])
		if result {
			return true, append(keys, words[0])
		}
		return false, nil

	}
	return false, nil

}
