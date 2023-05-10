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
	newElem := new(ListItem)
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
	newElem := new(ListItem)
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
		switch {
		case l.len == 1 && l.head == i:
			l.head = nil
			l.tail = nil
		case l.head == i:
			l.head = i.Next
			l.head.Prev = nil
		case l.tail == i:
			l.tail = i.Prev
			l.tail.Next = nil
		default:
			prev := i.Prev
			next := i.Next
			prev.Next = next
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
