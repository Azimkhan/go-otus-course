package week3_find_max

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortInts(t *testing.T) {
	// сортируем целые числа
	ints := make([]interface{}, 0, 5)
	ints = append(ints, 1, 10, 2, 3, 8)
	compareInts := func(i int, j int, data []interface{}) int {
		return data[i].(int) - data[j].(int)
	}

	result := FindMax(ints, compareInts)

	assert.Equal(t, 10, result)
}

func TestSortStrings(t *testing.T) {
	// сортируем строки по длине и алфавиту
	data := make([]interface{}, 0, 4)
	data = append(data, "John", "Bob", "Tom", "Samanta")
	compareStrings := func(i int, j int, data []interface{}) int {
		s1 := data[i].(string)
		s2 := data[j].(string)
		lenDiff := len(s1) - len(s2)
		if lenDiff != 0 {
			return lenDiff
		}

		r1 := []rune(s1)
		r2 := []rune(s2)
		for i := 0; i < len(r1); i++ {
			diff := r1[i] - r2[i]
			if diff != 0 {
				if diff > 0 {
					return -1
				}
				return 1
			}
		}
		return 0
	}

	result := FindMax(data, compareStrings)

	assert.Equal(t, "Samanta", result)
}
