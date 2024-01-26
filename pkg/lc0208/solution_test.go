package lc0208

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	t.Parallel()
	t.Run(`test_001`, func(t *testing.T) {
		trie := Constructor()
		trie.Insert("apple")
		trie.Insert("hotdog")
		assert.True(t, trie.Search("apple"))
		assert.False(t, trie.Search("app"))
		assert.True(t, trie.StartsWith("apple"))
		assert.False(t, trie.StartsWith("dog"))
	})
}
