package linked_list

import (
	"strconv"
)

type LinkedList[T any] struct {
	Size  uint
	first *node[T]
	last  *node[T]
}

type IndexOutOffBoundsErr struct {
	size  uint
	index uint
}

func (e IndexOutOffBoundsErr) Error() string {
	return "Collection has size: " + strconv.Itoa(int(e.size)) +
		" while requested index has value: " + strconv.Itoa(int(e.index))
}

func New[T any]() LinkedList[T] {
	return LinkedList[T]{
		Size:  0,
		first: nil,
		last:  nil,
	}
}

func (l *LinkedList[T]) Insert(item T) {
	newNode := node[T]{
		value: item,
		next:  nil,
	}

	if l.Size == 0 {
		l.first = &newNode
	} else {
		l.last.next = &newNode
		l.last = &newNode
	}

	l.last = &newNode
	l.Size++
}

func (l *LinkedList[T]) Get(index uint) (T, error) {
	if index >= l.Size {
		return *new(T), IndexOutOffBoundsErr{l.Size, index}
	}

	node := l.first
	for i := uint(0); i < index; i++ {
		node = node.next
	}

	return node.value, nil
}

func (l LinkedList[T]) Delete(index uint) error {
	if index >= l.Size {
		return IndexOutOffBoundsErr{l.Size, index}
	}

	previousNode := l.first
	for i := uint(0); i < index-1; i++ {
		previousNode = previousNode.next
	}

	previousNode.next = previousNode.next.next

	return nil
}

type node[T any] struct {
	value T
	next  *node[T]
}
