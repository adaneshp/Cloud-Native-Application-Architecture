// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func topWords(path string, K int) []WordCount {
	f, err := os.ReadFile(path)
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

	if K > len(wordCounts) {
		fmt.Println("K out of range. Please enter a number between 0 and", len(wordCounts), "and run the program again.")
	}
	return wordCounts[:K]
}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
