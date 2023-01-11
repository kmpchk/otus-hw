package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushBack(v interface{}) *ListItem {
	var newElem *ListItem = new(ListItem)
	newElem.Value = v
	newElem.Next = nil
	newElem.Prev = l.tail

	if newElem.Prev != nil {
		l.tail.Next = newElem
	} else {
		l.head = newElem
	}
	l.tail = newElem

	l.len++

	return newElem
}

func (l *list) PushFront(v interface{}) *ListItem {
	var newElem *ListItem = new(ListItem)
	newElem.Value = v
	newElem.Prev = nil
	newElem.Next = l.head

	if newElem.Next != nil {
		l.head.Prev = newElem
	} else {
		l.tail = newElem
	}
	l.head = newElem

	l.len++

	return newElem
}

func (l *list) Remove(i *ListItem) {
	if i != nil {
		prev := i.Prev
		next := i.Next

		if prev != nil {
			prev.Next = next
		}
		if next != nil {
			next.Prev = prev
		}

		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i != nil {
		if i != l.head {
			l.PushFront(i.Value)
			l.Remove(i)
		}
	}
}

func NewList() List {
	return new(list)
}
