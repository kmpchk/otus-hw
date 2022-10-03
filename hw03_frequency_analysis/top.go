package hw03frequencyanalysis

//package main

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

	var wordsFreqMap map[string]int = make(map[string]int)

	for _, word := range words {
		wordsFreqMap[word]++
	}

	//fmt.Println(wordsFreqMap)

	var wordsFreq []WordsFreq = make([]WordsFreq, 0)

	//fmt.Println(wordsFreq)

	for key, value := range wordsFreqMap {
		wordsFreq = append(wordsFreq, WordsFreq{key, value})
	}

	/*for _, wordFreq := range wordsFreq {
		fmt.Printf("Key: %s, Value: %d\n", wordFreq.Key, wordFreq.Value)
	}*/

	//fmt.Println("\nSort...\n")

	sort.Slice(wordsFreq, func(i, j int) bool {
		if wordsFreq[i].Value == wordsFreq[j].Value {
			return wordsFreq[i].Key < wordsFreq[j].Key
		}
		return wordsFreq[i].Value > wordsFreq[j].Value
	})

	/*for _, wordFreq := range wordsFreq {
		fmt.Printf("Key: %s, Value: %d\n", wordFreq.Key, wordFreq.Value)
	}*/

	var top10words []string = make([]string, 0)
	for _, wordFreq := range wordsFreq {
		top10words = append(top10words, wordFreq.Key)
	}

	//fmt.Println(top10words)

	if len(top10words) >= 10 {
		return top10words[:10]
	}
	return top10words
}
