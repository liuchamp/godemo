package main

type Trie struct {
	children map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.children = make(map[byte]*Trie)
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	s := word[0]
	if this.children == nil {
		this.children = make(map[byte]*Trie)
	}
	if c, ok := this.children[s]; ok {
		if c.children !=nil {
			return
		}
	} else {
		t := Constructor()
		this.children[s] = &t
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

	return false
}



/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
