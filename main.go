package main

import "fmt"

type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}
type List struct {
	storage []Item
}

func (s *List) Len() int {
	return len(s.storage)
}

func (s *List) First() *Item {
	return &s.storage[0]
}

func (s *List) Last() *Item {
	return &s.storage[len(s.storage)-1]
}

func (s *List) PushBack(v interface{}) {
	var newItem Item
	newItem.prev = s.Last()
	newItem.value = v
	s.Last().next = &newItem
	s.storage = append(s.storage, newItem)
}

func (s *List) PushFront(v interface{}) {
	var newItem Item
	newItem.next = s.First()
	newItem.value = v
	s.First().prev = &newItem
	s.storage = append([]Item{newItem}, s.storage...)
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

func (s *List) Find(x Item) int {
	for i, n := range s.storage {
		if x == n {
			return i
		}
	}
	return len(s.storage)
}

func (s *List) Remove(i Item) {
	itemInd := s.Find(i)
	s.storage = append(s.storage[:itemInd], s.storage[itemInd+1:]...)
	s.storage[itemInd].prev = &s.storage[itemInd-1]
	s.storage[itemInd].next = &s.storage[itemInd+1]
}

func main() {
	var itemList List
	var firstItem Item
	firstItem.value = 0
	itemList.storage = append(itemList.storage, firstItem)

	itemList.PushFront(123)
	itemList.PushBack(221)
	itemList.PushBack(224)
	fmt.Println(itemList.First())
	fmt.Println(itemList.Last())
	fmt.Println(itemList.Len())
	fmt.Println(itemList.storage[1].Next())
	fmt.Println(itemList.storage[1].Prev())
	fmt.Println(itemList.storage)
	itemList.Remove(itemList.storage[1])
	fmt.Println(itemList.storage)
}
