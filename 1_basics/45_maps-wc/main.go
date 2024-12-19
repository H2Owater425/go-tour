package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)

	var lastIndex int

	for i := range s {
		if s[i] == ' ' {
			wordCounts[s[lastIndex:i]]++

			lastIndex = i + 1
		}
	}

	wordCounts[s[lastIndex:]]++

	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
