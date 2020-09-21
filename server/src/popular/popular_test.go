package popular

import (
	"testing"
	"time"
)

func TestPopular(t *testing.T) {
	// Init this module
	Init()

	// Test when there is no message
	expectedMostPopularWord := ""
	gotMostPopularWord := mostPopularWord
	if expectedMostPopularWord != gotMostPopularWord {
		t.Errorf("Expected to get %s. But now get %s", expectedMostPopularWord, gotMostPopularWord)
		return
	}

	expectedMessageCount := 0
	gotMessageCount := getMessageCount()
	if expectedMessageCount != gotMessageCount {
		t.Errorf("Expected to get %d messages. But now we got %d.", expectedMessageCount, gotMessageCount)
		return
	}

	// Test after adding 5 messages.
	messageList := []string{
		"This is an apple.",
		"This is an red apple.",
		"I like orange and pear.",
		"What do you like?",
		"She doesn't like salad.",
	}

	for _, message := range messageList {
		AddMessage(message)
	}

	// Wait for 2 seconds and let the goroutine running in the background to calculate the most popular word
	time.Sleep(2 * time.Second)

	expectedMostPopularWord = "like"
	gotMostPopularWord = mostPopularWord
	if expectedMostPopularWord != gotMostPopularWord {
		t.Errorf("Expected to get %s. But now get %s", expectedMostPopularWord, gotMostPopularWord)
		return
	}

	expectedMessageCount = 5
	gotMessageCount = getMessageCount()
	if expectedMessageCount != gotMessageCount {
		t.Errorf("Expected to get %d messages. But now we got %d.", expectedMessageCount, gotMessageCount)
		return
	}

	// Wait for 5 seconds to let the old messages expire
	time.Sleep(conMessageMaxValidSeconds * time.Second)

	// Test after all the old messages are expired and add a new message.
	AddMessage("This is a new message which contains duplicate message.")

	// Wait for 2 seconds and let the goroutine running in the background to calculate the most popular word
	time.Sleep(2 * time.Second)

	expectedMostPopularWord = "message"
	gotMostPopularWord = mostPopularWord
	if expectedMostPopularWord != gotMostPopularWord {
		t.Errorf("Expected to get %s. But now get %s", expectedMostPopularWord, gotMostPopularWord)
		return
	}

	expectedMessageCount = 1
	gotMessageCount = getMessageCount()
	if expectedMessageCount != gotMessageCount {
		t.Errorf("Expected to get %d messages. But now we got %d.", expectedMessageCount, gotMessageCount)
		return
	}
}
