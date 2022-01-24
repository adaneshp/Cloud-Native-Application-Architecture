package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type WordCount struct {
	Word  string
	Count int
}

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func main() {
	fmt.Println("Please enter K for seeing top K words: ")
	var K int
	fmt.Scanln(&K)
	f, err := os.ReadFile("passage")
	check(err)
	s := string(f)
	words := strings.Fields(s)
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		if _, found := wordsMap[words[i]]; found {
			wordsMap[words[i]]++
		} else {
			wordsMap[words[i]] = 1
		}
	}
	wordCounts := make([]WordCount, len(wordsMap))
	j := 0
	for key, val := range wordsMap {
		wordCounts[j].Word = key
		wordCounts[j].Count = val
		j++
	}
	sortWordCounts(wordCounts)

	for K > len(wordCounts) {
		fmt.Println("K out of range. Please enter a number between 0 and", len(wordCounts))
		fmt.Scanln(&K)
	}
	fmt.Println(wordCounts[:K])
}
