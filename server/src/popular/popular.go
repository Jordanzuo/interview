/*Package popular ...
There are two ways to calculate the most popular word in the last 5 seconds:
1. Calculate whenever there is a request. If the frequency of this action is not high, it's acceptable.
2. Calculate it periodically and cache it for 1 or 2 seconds. This is not as real time as the first solution. But it can promote performance.
Here, I assume this action is called very frequently and the performance is more important than accuracy.
*/
package popular

import (
	"fmt"
	"sync"
	"time"
)

var (
	mostPopularWord              = ""
	messageListInLastFiveSeconds = make([]*Message, 0, 64)
	rwmutex                      sync.RWMutex
)

// Init ... Start a new goroutine to calculate the most popular word periodically
func Init() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("There is an unhandled error:%v", r)
			}
		}()

		for {
			time.Sleep(time.Second)
			calcMostPopularWord()
		}
	}()
}

func calcMostPopularWord() {
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	if len(messageListInLastFiveSeconds) == 0 {
		return
	}

	wordMap := make(map[string]int, 1024)
	for _, v := range messageListInLastFiveSeconds {
		for _, token := range v.tokenize() {
			wordMap[token]++
		}
	}

	word, count := "", 0
	for k, v := range wordMap {
		if v > count {
			word, count = k, v
		}
	}

	mostPopularWord = word
}

func getMessageCount() int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return len(messageListInLastFiveSeconds)
}

// AddMessage ... Add new message to message list for calculating the most popular word
// message: new message
func AddMessage(message string) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	messageListInLastFiveSeconds = append(messageListInLastFiveSeconds, newMessage(message))

	// Delete expired messages
	// Since all the messages are stored orderly by SendTime
	// So we can iterate the list and find the first one which is not expired and termincate
	validIndex := len(messageListInLastFiveSeconds) - 1
	for i, v := range messageListInLastFiveSeconds {
		if v.expired() == false {
			validIndex = i
			break
		}
	}

	// If there are some expired messages, remove them
	if validIndex > 0 {
		messageListInLastFiveSeconds = messageListInLastFiveSeconds[validIndex:]
	}
}
