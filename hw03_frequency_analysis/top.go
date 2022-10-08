package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordsFreq struct {
	Key   string
	Value int
}

func Top10(inputString string) []string {
	words := strings.Fields(inputString)
	if len(words) == 0 {
		return make([]string, 0)
	}

	wordsFreqMap := make(map[string]int, 0)

	for _, word := range words {
		wordsFreqMap[word]++
	}

	wordsFreq := make([]WordsFreq, 0)

	for key, value := range wordsFreqMap {
		wordsFreq = append(wordsFreq, WordsFreq{key, value})
	}

	sort.Slice(wordsFreq, func(i, j int) bool {
		if wordsFreq[i].Value == wordsFreq[j].Value {
			return wordsFreq[i].Key < wordsFreq[j].Key
		}
		return wordsFreq[i].Value > wordsFreq[j].Value
	})

	top10words := make([]string, 0)
	for _, wordFreq := range wordsFreq {
		top10words = append(top10words, wordFreq.Key)
	}

	if len(top10words) >= 10 {
		return top10words[:10]
	}

	return top10words
}
