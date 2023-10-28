package main

import (
	"reflect"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	tests := []struct {
		letters  string
		expected []string
	}{
		{"abc", []string{"", "a", "ab", "abc", "ac", "acb", "b", "ba", "bac", "bc", "bca", "c", "ca", "cab", "cb", "cba"}},
		{"xy", []string{"", "x", "xy", "y", "yx"}},
		{"", []string{""}},
	}

	for _, test := range tests {
		letters := []rune(test.letters)
		var permutations []string
		GeneratePermutations(letters, "", &permutations)

		if !reflect.DeepEqual(permutations, test.expected) {
			t.Errorf("For input '%s', expected permutations: %v, got: %v", test.letters, test.expected, permutations)
		}
	}
}

func TestGeneratePermutationsWithLength(t *testing.T) {
	tests := []struct {
		letters  string
		length   int
		expected []string
	}{
		{"abc", 1, []string{"a", "b", "c"}},
		{"xy", 2, []string{"xy", "yx"}},
		{"", 0, []string{""}},
	}

	for _, test := range tests {
		letters := []rune(test.letters)
		var permutations []string
		GeneratePermutationsWithLength(test.length, letters, "", &permutations)

		if !reflect.DeepEqual(permutations, test.expected) {
			t.Errorf("For input '%s', expected permutations: %v, got: %v", test.letters, test.expected, permutations)
		}
	}
}

func TestTrieInsertAndSearch(t *testing.T) {
	trie := NewTrie()

	// Insert words into the Trie
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("carrot")
	// Test the Search method
	tests := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"banana", true},
		{"carrot", true},
		{"grape", false},
		{"car", false},
		{"apples", false},
	}

	for _, test := range tests {
		found := trie.Search(test.word)
		if found != test.expected {
			t.Errorf("Expected %v for word '%s', but got %v", test.expected, test.word, found)
		}
	}
}

func TestTrieSearchEmptyTrie(t *testing.T) {
	trie := NewTrie()

	// Test searching in an empty Trie
	tests := []struct {
		word     string
		expected bool
	}{
		{"apple", false},
		{"banana", false},
	}

	for _, test := range tests {
		found := trie.Search(test.word)
		if found != test.expected {
			t.Errorf("Expected %v for word '%s', but got %v", test.expected, test.word, found)
		}
	}
}

func TestTrieSearchPartialWord(t *testing.T) {
	trie := NewTrie()

	// Insert words into the Trie
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("carrot")

	// Test searching for partial words
	tests := []struct {
		word     string
		expected bool
	}{
		{"app", false},
		{"ban", false},
		{"car", false},
	}

	for _, test := range tests {
		found := trie.Search(test.word)
		if found != test.expected {
			t.Errorf("Expected %v for partial word '%s', but got %v", test.expected, test.word, found)
		}
	}
}
