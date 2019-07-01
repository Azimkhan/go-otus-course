package week2

import (
	"testing"
	"io/ioutil"
	"strings"
	"gotest.tools/assert"
	"fmt"
)

func TestTop10WordsByFrequency(t *testing.T) {
	bytes, err := ioutil.ReadFile("sample_text.txt")
	if err != nil {
		t.Fail()
	}
	// убираем переносы строк из текста
	text := strings.Replace(string(bytes), "\n", " ", -1)
	frequencies := Top10WordsByFrequency(text)
	for i,f := range frequencies {
		fmt.Printf("%d. %v: %v\n", i+1, f.word, f.count)
	}
	assert.Equal(t,10, len(frequencies))
	assert.Equal(t, "go", frequencies[0].word)
}