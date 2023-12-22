package main

/*
型の制約を持たせた状態でSetが作成可能
*/

// Set - 汎用型Tを用いたSetの実装
type Set[T comparable] struct {
	items map[T]struct{}
}

// New - 新しいSetを作成
func New[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

// Add - 要素をSetに追加
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Remove - 要素をSetから削除
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains - 要素がSetに含まれているかをチェック
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Items - Setの要素を取得
func (s *Set[T]) Items() []T {
	var items []T
	for item := range s.items {
		items = append(items, item)
	}
	return items
}
