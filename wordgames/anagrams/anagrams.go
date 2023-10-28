package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

var (
	dict    string
	letters string
)

func main() {
	// Define command-line arguments and options
	flag.StringVar(&dict, "dict", "", "Dictionary for search")
	flag.StringVar(&letters, "letters", "", "Given letters for search")

	// Parse the command-line arguments
	flag.Parse()

	// Check if required arguments are provided
	if dict == "" || letters == "" {
		fmt.Println("Error: Argument 'dict' or 'letters' is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	trie, err := buildTrieFromFile(dict)
	if err != nil {
		fmt.Printf("Error reading dictionary and building Trie: %v\n", err)
		return
	}

	lettersRune := []rune(letters)

	wordsToSearch := []string{}
	GeneratePermutations(lettersRune, "", &wordsToSearch)
	// Test the Trie with some example words
	for _, word := range wordsToSearch {
		found := trie.Search(word)
		fmt.Printf("Word '%s' found in dictionary: %v\n", word, found)
	}
}

func buildTrieFromFile(filename string) (*Trie, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	trie := NewTrie()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		trie.Insert(word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return trie, nil
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}
func GeneratePermutations(letters []rune, current string, permutations *[]string) {
	for i := 0; i <= len(letters); i++ {
		GeneratePermutationsWithLength(i, letters, "", permutations)
	}
	sort.Strings(*permutations)
}

func GeneratePermutationsWithLength(length int, letters []rune, current string, permutations *[]string) {
	if len(current) == length {
		*permutations = append(*permutations, current)
		return
	}

	for i, letter := range letters {
		// Create a copy of the letters without the current letter
		remainingLetters := make([]rune, len(letters))
		copy(remainingLetters, letters)
		remainingLetters = append(remainingLetters[:i], remainingLetters[i+1:]...)

		// Recursively generate permutations
		GeneratePermutationsWithLength(length, remainingLetters, current+string(letter), permutations)
	}
}
