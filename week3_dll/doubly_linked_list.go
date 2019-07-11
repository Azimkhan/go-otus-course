package week3_dll

type List struct {
	first *Item
	last  *Item
}

func (l *List) Len() (len int) {
	i := l.first

	for i != nil {
		len++
		i = i.next
	}
	return
}

func (l *List) First() *Item {
	return l.first
}

func (l *List) Last() *Item {
	return l.last
}

func (l *List) PushFront(v interface{}) {

	newNode := &Item{
		list:  l,
		value: v,
	}
	if l.first == nil {

		l.first = newNode
		l.last = newNode
		return
	}
	l.insertBefore(l.first, newNode)

}
func (l *List) PushBack(v interface{}) {
	newNode := &Item{
		list:  l,
		value: v,
	}
	if l.last == nil {
		l.first = newNode
		l.last = newNode
		return
	}
	l.insertAfter(l.last, newNode)
}
func (l *List) insertBefore(node *Item, newNode *Item) {
	newNode.next = node
	if node.prev == nil {
		newNode.prev = nil
		l.first = newNode
	} else {
		node.prev.next = newNode
		newNode.prev = node.prev
	}
	node.prev = newNode
}
func (l *List) insertAfter(node *Item, newNode *Item) {
	newNode.prev = node
	if node.next == nil {
		newNode.next = nil
		l.last = newNode
	} else {
		newNode.next = node.next
		node.next.prev = newNode
	}
	node.next = newNode
}

type Item struct {
	list  *List
	next  *Item
	prev  *Item
	value interface{}
}

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Prev() *Item {
	return i.prev
}

func (i *Item) Remove() {
	// next && prev == nil
	if i.next == nil && i.prev == nil {
		i.list.first = nil
		i.list.last = nil
		return
	}
	// next == nil
	if i.next == nil {
		i.list.last = i.prev
		i.prev.next = nil
		return
	} else if i.prev == nil {
		// prev == nil
		i.list.first = i.next
		i.next.prev = nil
		return
	}

	i.prev.next = i.next
	i.next.prev = i.prev

}
