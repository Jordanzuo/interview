package popular

import (
	"strings"
	"time"
)

const (
	conMessageMaxValidSeconds = 5
)

var (
	stopWordsMap    = make(map[string]struct{}, 8)
	punctuationList = make([]string, 0, 16)
)

func init() {
	stopWordsMap["a"] = struct{}{}
	stopWordsMap["an"] = struct{}{}
	stopWordsMap["the"] = struct{}{}
	stopWordsMap["yes"] = struct{}{}
	stopWordsMap["no"] = struct{}{}
	stopWordsMap["or"] = struct{}{}
	stopWordsMap["and"] = struct{}{}
	// There should be more stop words
}

func init() {
	punctuationList = append(punctuationList, ",")
	punctuationList = append(punctuationList, ";")
	punctuationList = append(punctuationList, ".")
	punctuationList = append(punctuationList, "?")
	punctuationList = append(punctuationList, "!")
	punctuationList = append(punctuationList, "'")
	punctuationList = append(punctuationList, "\"")
	punctuationList = append(punctuationList, "|")
	punctuationList = append(punctuationList, "[")
	punctuationList = append(punctuationList, "]")
	punctuationList = append(punctuationList, "{")
	punctuationList = append(punctuationList, "}")
}

// Message ... This represents the message to calculate the most popular word
type Message struct {
	SendTime int64
	Content  string
}

func (this *Message) tokenize() (tokens []string) {
	// Different language requires different tokenizer, but here just for simplicity, I just use space to split words
	words := strings.Split(this.Content, " ")

	// Stop words and punctuation should be removed
	for _, word := range words {
		if _, exists := stopWordsMap[word]; exists {
			continue
		}
		for _, puctuation := range punctuationList {
			word = strings.ReplaceAll(word, puctuation, "")
		}
		tokens = append(tokens, word)
	}

	return
}

func (this *Message) expired() bool {
	return this.SendTime < time.Now().Unix()-conMessageMaxValidSeconds
}

func newMessage(content string) *Message {
	return &Message{
		SendTime: time.Now().Unix(),
		Content:  content,
	}
}
