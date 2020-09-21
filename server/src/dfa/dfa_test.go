package dfa

import (
	"fmt"
	"testing"
)

func init() {
	err := Init("../../config/profanity_words.txt")
	if err != nil {
		panic(fmt.Sprintf("Init profanity words file failed. Error:%s", err))
	}
}

func TestIsMatch(t *testing.T) {
	input := "bitch, I don't want to be rude. bitch, Just for test. bitch"
	expected := true
	got := IsMatch(input)
	if expected != got {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}

	input = "I want to be polite."
	expected = false
	got = IsMatch(input)
	if expected != got {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}
}

func TestHandleWord(t *testing.T) {
	input := "bitch, I don't want to be rude. bitch, Just for test. bitch"
	expected := "*****, I don't want to be rude. *****, Just for test. *****"
	got := HandleWord(input, '*')
	if expected != got {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}

	input = "I want to be polite"
	expected = "I want to be polite"
	got = HandleWord(input, '*')
	if expected != got {
		t.Errorf("Expected to get %v, but got %v", expected, got)
	}
}

func BenchmarkIsMatch(b *testing.B) {
	input := `1. From the time that you receive this exercise document, you have ​ 24 hours​ to complete it
	and return the final deliverables.
	2. Your final delivery should be a git repository
	3. Take note of all NPM packages that you add, you will be expected to explain what they
	are and why you used them`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMatch(input)
	}
}
