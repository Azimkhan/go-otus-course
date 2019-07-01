package week2

import (
	"strings"
	"sort"
)

type WordFrequency struct {
	word  string
	count int
}

func Top10WordsByFrequency(text string) []WordFrequency {
	wordMap := map[string]int{}
	// создаем словарь слово -> частота
	words := strings.Split(strings.ToLower(text), " ")
	for _, word := range words {
		if len([]rune(word)) > 1 {
			wordMap[word] = wordMap[word] + 1
		}
	}

	// преорбразуем словарь в массив структур
	numWords := len(wordMap)
	pairs := make([]WordFrequency, numWords)
	for w, c := range wordMap {
		pairs = append(pairs, WordFrequency{w, c})
	}

	// сортируем массив по частоте
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// берем только топ
	if numWords < 10 {
		return pairs[:numWords]
	}
	return pairs[:10]
}
