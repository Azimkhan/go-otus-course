package task2_shortener

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestShorten(t *testing.T) {
	shortener := NewStandardShortener()
	urls := []string{"https://google.com", "https://yandex.ru", "https://mail.ru"}
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(originalUrl string, g *sync.WaitGroup) {
			defer g.Done()
			shortUrl := shortener.Shorten(originalUrl)
			fmt.Printf("%v -> %v\n", originalUrl, shortUrl)
			resolvedUrl := shortener.Resolve(shortUrl)
			assert.Equal(t, resolvedUrl, originalUrl)
		}(url, &wg)
	}

	wg.Wait()

}
