package week3

import (
	"gotest.tools/assert"
	"testing"
)

func TestList(t *testing.T) {
	list1 := List{}
	list1.PushFront("hello")
	assert.Equal(t, 1, list1.Len())
	assert.Equal(t, "hello", list1.First().Value())
	assert.Equal(t, "hello", list1.Last().Value())

	list1.PushFront("foo")
	assert.Equal(t, 2, list1.Len())
	assert.Equal(t, "foo", list1.First().Value())
	assert.Equal(t, "hello", list1.Last().Value())

	list1.PushFront("baz")
	assert.Equal(t, 3, list1.Len())
	assert.Equal(t, "baz", list1.First().Value())

	// remove first element
	list1.first.Remove()
	assert.Equal(t, 2, list1.Len())
	assert.Equal(t, "foo", list1.First().Value())

	list2 := List{}
	list2.PushBack("hello")
	assert.Equal(t, 1, list2.Len())
	assert.Equal(t, "hello", list2.First().Value())
	assert.Equal(t, "hello", list2.Last().Value())

	list2.PushBack("foo")
	assert.Equal(t, 2, list2.Len())
	assert.Equal(t, "hello", list2.First().Value())
	assert.Equal(t, "foo", list2.Last().Value())

	list2.PushBack("baz")
	assert.Equal(t, 3, list2.Len())
	assert.Equal(t, "hello", list2.First().Value())
	assert.Equal(t, "baz", list2.Last().Value())

	// remove element from middle of list
	list2.First().Next().Remove()
	assert.Equal(t, 2, list1.Len())
	assert.Equal(t, "hello", list2.First().Value())
	assert.Equal(t, "baz", list2.Last().Value())

}
