package task2_shortener

import (
	"sync/atomic"
	"fmt"
	"strings"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type StandardShortener struct {
	baseUrl string
	counter uint64
	mem     map[string]string
}

func NewStandardShortener() *StandardShortener {
	s := new(StandardShortener)
	s.baseUrl = "https://short.link/"
	s.mem = make(map[string]string)
	s.counter = 0
	return s
}
func (s *StandardShortener) Resolve(url string) string {
	if !strings.HasPrefix(url, s.baseUrl) {
		return url
	}

	id := url[len(s.baseUrl):]
	original, ok := s.mem[id]
	if !ok {
		return url
	} else {
		return original
	}
}

func (s *StandardShortener) Shorten(url string) string {
	// в качестве генератора уникального идентификатора ссылки используем
	// простой счетчик, используем atomic чтобы код был потокобезопасным
	currentNum := atomic.AddUint64(&s.counter, 1)
	id := fmt.Sprintf("%x", currentNum)
	newUrl := s.baseUrl + id
	s.mem[id] = url
	return newUrl
}
