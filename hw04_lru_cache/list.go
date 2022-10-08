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

	if (l.head == nil) && (l.tail == nil) {
		newElem.Prev = nil
		newElem.Next = nil
		l.head = newElem
		l.tail = newElem
	} else {
		//var exch_var *ListItem = l.tail
		//newElem.Prev = l.head
		l.Back().Next = newElem
		//l.head = l.tail
		l.tail = newElem
		//l.head.Next = l.tail
		newElem.Prev = l.tail
		newElem.Next = nil
	}

	l.len++

	return newElem
}

func (l *list) PushFront(v interface{}) *ListItem {
	var newElem *ListItem = new(ListItem)
	newElem.Value = v

	if (l.head == nil) && (l.tail == nil) {
		newElem.Prev = nil
		newElem.Next = nil
		l.head = newElem
		l.tail = newElem
	} else {
		//var exch_var *ListItem = l.tail
		//newElem.Prev = l.head
		//l.tail = l.head
		l.Front().Prev = newElem
		newElem.Next = l.head
		l.head = newElem
		//l.head.Next = l.tail
		newElem.Prev = nil
	}

	l.len++

	return newElem
}

func (l *list) Remove(i *ListItem) {
	if i != nil {
		prev := i.Prev
		next := i.Next

		if prev != nil {
			prev.Next = next
			//fmt.Println(next)
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
			//l.Front().Value = i.Value
			head := l.head
			l.head = i
			l.head.Prev = nil
			l.head.Next = head
			head.Prev = l.head
		}
	}
}

func NewList() List {
	return new(list)
}
