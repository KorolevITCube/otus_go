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
	first *ListItem
	last  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	curr := &ListItem{
		Value: v,
	}
	if l.first != nil {
		l.first.Prev = curr
		curr.Next = l.first
	} else {
		l.last = curr
	}
	l.first = curr
	l.len++
	return curr
}

func (l *list) PushBack(v interface{}) *ListItem {
	curr := &ListItem{
		Value: v,
	}
	if l.last != nil {
		l.last.Next = curr
		curr.Prev = l.last
	} else {
		l.first = curr
	}
	l.last = curr
	l.len++
	return curr
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.first != i && l.first != nil {
		// удаляем из текущего места.
		if i.Prev != nil {
			i.Prev.Next = i.Next
		}
		if i.Next != nil {
			i.Next.Prev = i.Prev
		}
		if l.last == i {
			l.last = i.Prev
		}
		// обновляем ссылки.
		i.Prev = nil
		i.Next = l.first
		// ставим в начало.
		l.first.Prev = i
		l.first = i
	}
}

func NewList() List {
	return new(list)
}
