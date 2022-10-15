package langdet

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func requireNoError(t interface{ Fatal(a ...any) }, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertFalse(t interface{ Error(a ...any) }, expr bool, msg string) {
	if expr {
		t.Error(msg)
	}
}

func TestDetectWikipediaPage(t *testing.T) {
	internet, err := readFile("testfiles/internet_eng.txt")
	requireNoError(t, err)
	res := DetectLanguage(internet)
	fmt.Println(res)
	assertFalse(t, len(res) == 0 || res[0].Tag != language.English, "not english")
}

func TestDetectShortSentence(t *testing.T) {
	test := "A sentence is a group of words put together in a complete, meaningful way. It expresses a thought, statement, question, wish, command, suggestion, or idea. We use sentences every day when we’re writing and speaking. What’s more, a sentence combines words in a grammatically correct way. There are lots things to understand and of rules to follow when making a sentence, from punctuation, to word order, to making sure you have all of the right parts. This article will cover everything you need to know about strong sentences!"
	res := DetectLanguage([]byte(test))
	fmt.Println(res)
	assertFalse(t, len(res) == 0 || res[0].Tag != language.English, "not english")
}

func BenchmarkDetectLanguage(b *testing.B) {
	internet, err := readFile("testfiles/internet_eng.txt")
	requireNoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DetectLanguage(internet)
	}
}
