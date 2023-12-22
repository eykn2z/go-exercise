package main

type Set struct {
	items map[interface{}]bool
}

func New() *Set {
	return &Set{items: make(map[interface{}]bool)}
}

func (s *Set) Add(item interface{}) {
	s.items[item] = true
}

func (s *Set) Remove(item interface{}) {
	delete(s.items, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Items() []interface{} {
	var items []interface{}
	for item := range s.items {
		items = append(items, item)
	}
	return items
}
