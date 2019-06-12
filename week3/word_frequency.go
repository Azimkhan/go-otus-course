package week3

import (
	"strings"
	"sort"
)

type WordFrequency struct {
	word  string
	count int
}

func Top10WordsByFrequency(text string) []WordFrequency {
	out := map[string]int{}
	// создаем словарь слово -> частота
	words := strings.Split(strings.ToLower(text), " ")
	for _, word := range words {
		if len([]rune(word)) > 1 {
			out[word] = out[word] + 1
		}
	}

	// преорбразуем словарь в массив структур
	numWords := len(out)
	pairs := make([]WordFrequency, numWords)
	i := 0
	for w, c := range out {
		pairs[i] = WordFrequency{w, c}
		i++
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
