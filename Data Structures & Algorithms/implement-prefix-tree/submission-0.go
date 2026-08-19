type TrieNode struct {
isWord bool
children map[rune]*TrieNode
}

type PrefixTree struct {
    root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{
        root: &TrieNode{
            children: make(map[rune]*TrieNode),
        },
    }
}

func (t *PrefixTree) Insert(word string) {
curr := t.root
for _, c := range word {
    if _, ok := curr.children[c]; !ok {
        curr.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
    }
    curr = curr.children[c]
}
curr.isWord = true
}

func (t *PrefixTree) Search(word string) bool {
curr := t.root
for _, c := range word {
    if _, ok := curr.children[c]; !ok {
        return false
    }
    curr = curr.children[c]
}
return curr.isWord
}

func (t *PrefixTree) StartsWith(prefix string) bool {
curr := t.root
for _, c := range prefix {
    if _, ok := curr.children[c]; !ok {
        return false
    }
    curr = curr.children[c]
}
return true
}
