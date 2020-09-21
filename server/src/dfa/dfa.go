/*
See Profanity Filter design.md for more information
*/
package dfa

import (
	"io/ioutil"
	"strings"
)

var (
	dfaObj *dfa
)

// dfa object
type dfa struct {
	// The root node
	root *trieNode
}

// Since go doesn't have tuple type. In order to define extra struct, I use two []int to represent start index and end index pairs.
// When handling return values from this function, the two list should be used accordingly.
func (this *dfa) searcSentence(sentence string) (startIndexList, endIndexList []int) {
	// Point current node to the root node and initialize some variables.
	currNode := this.root
	start, end, valid := 0, 0, false

	// Iterate the setence to handle each letter
	sentenceRuneList := []rune(sentence)
	for i := 0; i < len(sentenceRuneList); {
		// If the letter can be found in current node's children, then continue to find along this path.
		if child, exists := currNode.children[sentenceRuneList[i]]; exists {
			// If the letter is end of a word, then it's a valid match.
			// Then set valid to true, and assign the index to the end variable.
			if child.isEndOfWord {
				end = i
				valid = true
			}

			// If the child doesn't have any child, it means it's the end of a path. Then add the last valid index pair into list.
			// And continue to handle the next letter from the root node.
			// Otherwise, continue to handle along this path.
			if len(child.children) == 0 {
				startIndexList = append(startIndexList, start)
				endIndexList = append(endIndexList, end)
				currNode = this.root

				// Reset variables, and starts from the next letter.
				start, end, valid = i+1, 0, false
			} else {
				currNode = child
			}

			// Handle the next letter.
			i++
		} else {
			// When the letter can't be found in current node's children, there are two possibilities:
			// 1. There is already a valid match index pair. Then add them to list. And rehandle this letter again from the root node.
			// 2. There is no valid match index pair. Then continue to handle next letter from the root node.
			if valid {
				startIndexList = append(startIndexList, start)
				endIndexList = append(endIndexList, end)
				currNode = this.root
				start, end, valid = i, 0, false
			} else {
				currNode = this.root
				start, end, valid = i+1, 0, false
				i++
			}
		}
	}

	// Find if there is any valid pairs which hasn't been processed.
	if valid {
		startIndexList = append(startIndexList, start)
		endIndexList = append(endIndexList, end)
	}

	return
}

// insert new word into object
func (this *dfa) insertWord(word []rune) {
	currNode := this.root
	for _, c := range word {
		if cildNode, exist := currNode.children[c]; !exist {
			cildNode = newTrieNode()
			currNode.children[c] = cildNode
			currNode = cildNode
		} else {
			currNode = cildNode
		}
	}

	currNode.isEndOfWord = true
}

// Check if there is any word in the trie that starts with the given prefix.
func (this *dfa) startsWith(prefix []rune) bool {
	currNode := this.root
	for _, c := range prefix {
		if cildNode, exist := currNode.children[c]; !exist {
			return false
		} else {
			currNode = cildNode
		}
	}

	return true
}

// Init ... Init dfa file and object
// filePath: The dfa content file path
// Return values:
// Error if exists
func Init(filePath string) (err error) {
	var dfaContent []byte
	dfaContent, err = ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	dfaObj = &dfa{
		root: newTrieNode(),
	}

	wordList := strings.Split(string(dfaContent), "\n")
	for _, word := range wordList {
		wordRuneList := []rune(word)
		if len(wordRuneList) > 0 {
			dfaObj.insertWord(wordRuneList)
		}
	}

	return
}

// IsMatch ... Judge if input sentence contains some special caracter
// Return:
// Matc or not
func IsMatch(sentence string) bool {
	startIndexList, _ := dfaObj.searcSentence(sentence)
	return len(startIndexList) > 0
}

// HandleWord ... Handle sentence. Use specified caracter to replace those sensitive caracters.
// input: Input sentence
// replaceCh: candidate
// Return:
// Sentence after manipulation
func HandleWord(sentence string, replaceCh rune) string {
	startIndexList, endIndexList := dfaObj.searcSentence(sentence)
	if len(startIndexList) == 0 {
		return sentence
	}

	// Manipulate
	sentenceList := []rune(sentence)
	for i := 0; i < len(startIndexList); i++ {
		for index := startIndexList[i]; index <= endIndexList[i]; index++ {
			sentenceList[index] = replaceCh
		}
	}

	return string(sentenceList)
}
