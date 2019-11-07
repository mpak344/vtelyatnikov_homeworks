package DoublyLinkedList

type Item struct {
	previous *Item
	next *Item
	value interface{}
}

// возвращает значение Item-a
func (item Item) Value() interface{} {
	return item.value
}

// возвращает следующий Item
func (item Item) Nex() *Item {
	return item.next
}

// возвращает предыдущий Item
func (item Item) Prev() *Item {
	return item.previous
}

type List struct {
	first *Item
	last *Item
	len int
}

// возвращает длину списка
func (list *List) Len() int {
	return list.len
}

// возвращает первый элемент списка
func (list *List) First() *Item {
	return list.first
}

// возвращает последний элемент списка
func (list *List) Last() *Item {
	return list.last
}

// добавляет значение в начало
func (list *List) PushFront(v interface{}) {
	currentItem := new(Item)
	currentItem.value = v

	previousItem := list.first
	if (previousItem != nil) {
		previousItem.previous = currentItem
		currentItem.next = previousItem
	}
	if (list.last == nil) {
		 list.last = currentItem 
	}	
	list.first = currentItem
	list.len++
}

// добавляет значение в конец
func (list *List) PushBack(v interface{}) {
	currentItem := new(Item)
	currentItem.value = v

	nextItem := list.last
	if (nextItem != nil) {
		nextItem.next = currentItem
		currentItem.previous = nextItem
	}
	if (list.first == nil) {
		list.first = currentItem
	}
	list.last = currentItem
	list.len++
}

// удаляет Item из списка 
func (list *List) Remove(i Item) {
	previous := i.Prev()
	next := i.Nex()
	if (previous != nil) {
		previous.next = next
	} else {
		list.first = next
	}

	if (next != nil) {
		next.previous = previous
	} else {
		list.last = previous
	}

	list.len--
}