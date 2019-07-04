package week3

import (
	"testing"
	"gotest.tools/assert"
)

func TestList(t *testing.T) {
	list := List{}
	list.PushFront("hello")
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, "hello", list.First().Value())
	assert.Equal(t, "hello", list.Last().Value())

	list.PushFront("foo")
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, "foo", list.First().Value())
	assert.Equal(t, "hello", list.Last().Value())
}
