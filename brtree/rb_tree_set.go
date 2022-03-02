package brtree

import (
	"data_structures_poligon/utils"
)

type leafColor = uint

const (
	RED leafColor = iota
	BLACK
)

type RBTreeSet[T any, C comparator.Compare[T]] struct {
	root        *treeNode[T]
	compare     comparator.Compare[T]
	blackHeight uint
}

type treeNode[T any] struct {
	value T
	color leafColor
	left  *treeNode[T]
	right *treeNode[T]
}

func NewPrimitiveSet[T any]() RBTreeSet[T] {
	return RBTreeSet[T]{
		root:        nil,
		blackHeight: uint(0),
	}
}

func (tree *RBTreeSet[T]) Insert(value T) {
	if tree.blackHeight == 0 {
		tree.root = &treeNode[T]{
			color: BLACK,
			value: value,
		}
		tree.blackHeight = 1
		return
	}

	node := tree.root
	for {
		i := tree.compare()
		if i < 0 {
			if node.left != nil {
				node = node.left
				continue
			}
			node.left = &treeNode[T]{
				value: value,
				color: RED,
			}
			break

		} else if i > 0 {
			if node.right != nil {
				node = node.right
				continue
			}
			node.right = &treeNode[T]{}
			break
		} else {
			return
		}
	}

}
