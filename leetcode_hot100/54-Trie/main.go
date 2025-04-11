package main

import "fmt"

func main() {
	word := "1234"
	prefix := "134"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search(word)
	param_3 := obj.StartsWith(prefix)
	fmt.Println(param_2)
	fmt.Println(param_3)
}

type Trie struct {
	l []string
}

func Constructor() Trie {
	return Trie{
		l: []string{},
	}
}

func (this *Trie) Insert(word string) {
	this.l = append(this.l, word)
}

func (this *Trie) Search(word string) bool {
	for i := 0; i < len(this.l); i++ {
		if this.l[i] == word {
			return true
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(this.l); i++ {
		if len(this.l[i]) < len(prefix) {
			continue
		}
		if this.l[i][:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
