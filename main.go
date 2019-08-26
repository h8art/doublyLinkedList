package main

import "fmt"

type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}
type List struct {
	FirstItem *Item
	LastItem  *Item
	len       int
}

func (s *List) PushBack(v interface{}) {
	var newItem Item
	newItem.prev = s.LastItem
	newItem.value = v
	s.LastItem.next = &newItem
	s.LastItem = &newItem
	s.len++
}

func (s *List) PushFront(v interface{}) {
	var newItem Item
	newItem.next = s.FirstItem
	newItem.value = v
	s.FirstItem.prev = &newItem
	s.FirstItem = &newItem
	s.len++
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

func (s *List) Find(i Item) *Item {
	item := s.FirstItem
	for j := 0; j < s.len; j++ {
		if *item == i {
			return item
		} else {
			item = item.next
		}
	}
	return nil
}

func (s *List) Remove(i Item) {
	item := s.FirstItem
	for j := 0; j < s.len; j++ {
		if i == *item {
			item.prev.next = item.next
			item.next.prev = item.prev
			s.len--
		} else {
			item = item.next
		}
	}
}

func main() {
	var itemList List
	var firstItem Item
	firstItem.value = 0
	itemList.FirstItem = &firstItem
	itemList.LastItem = &firstItem
	itemList.len = 1
	itemList.PushFront(123)
	itemList.PushBack(221)
	itemList.PushBack(224)
	fmt.Println(itemList.FirstItem)
	fmt.Println(itemList.LastItem)
	fmt.Println(itemList.len)
	var tempItem = *itemList.FirstItem.next
	fmt.Println(tempItem)
	fmt.Println("found item - ", itemList.Find(tempItem))
	itemList.Remove(tempItem)
}
