package pilha

import "fmt"

type Pilha[T any] struct {
	items   []T
	maxSize *int
	top     T
}

func Init[T any](maxSize *int) *Pilha[T] {
	return &Pilha[T]{
		items:   []T{},
		maxSize: maxSize,
	}
}

func (f *Pilha[T]) Push(v T) {
	f.items = append(f.items, v)
	f.top = v
}

func (f *Pilha[T]) Pop() (T, bool) {
	var zero T
	if len(f.items) == 0 {
		return zero, false
	}

	oldTop := f.top

	f.items = f.items[:len(f.items)-1]
	if len(f.items) == 0 {
		f.top = zero
		return oldTop, true
	}

	f.top = f.items[len(f.items)-1]
	return oldTop, true
}

func (f *Pilha[T]) List() {
	for _, item := range f.items {
		fmt.Println(item)
	}
}

func (f *Pilha[T]) IsEmpty() bool {
	return len(f.items) == 0
}

func (f *Pilha[T]) IsFull() bool {
	if f.maxSize == nil {
		return false
	}
	return len(f.items) == *f.maxSize
}

func (f *Pilha[T]) Size() int {
	return len(f.items)
}

func (f *Pilha[T]) Top() (T, bool) {
	if len(f.items) == 0 {
		var zero T
		return zero, false
	}
	return f.top, true
}

func (f *Pilha[T]) Items() []T {
	return append([]T(nil), f.items...)
}
