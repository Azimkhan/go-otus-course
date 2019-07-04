package week3

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
		value: v,
	}
	if l.first == nil {

		l.first = newNode
		l.last = newNode
		return
	}
	l.insertBefore(newNode, l.first)


}

func (l *List) insertAfter(newNode *Item, node *Item)  {
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

func (*List) PushBack(v interface{}) {
	panic("implement me")
}

type Item struct {
	next  *Item
	prev  *Item
	value interface{}
}

func (i *Item) Value() interface{} {
	return i.value
}

func (*Item) Next() *Item {
	panic("implement me")
}

func (*Item) Prev() *Item {
	panic("implement me")
}

func (*Item) Remove() {
	panic("implement me")
}
