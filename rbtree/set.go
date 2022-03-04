package rbtree

import (
	. "data_structures_poligon/utils"
)

type leafColor = uint8

const (
	RED leafColor = iota
	BLACK
)

type Set[T any, C Compare[T]] struct {
	root        *treeNode[T]
	compare     Compare[T]
	blackHeight uint
}

type treeNode[T any] struct {
	value  T
	color  leafColor
	parent *treeNode[T]
	left   *treeNode[T]
	right  *treeNode[T]
}

func newNode[T any](value T, parent *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{value: value, color: RED, parent: parent}
}

func NewPrimitiveSet[T any]() *Set[T] {
	return &Set[T]{
		root:        nil,
		blackHeight: uint(0),
	}
}

func NewSet[T any](comparator Compare[T]) *Set[T] {
	return &Set[T]{
		root:        nil,
		compare:     comparator,
		blackHeight: 0,
	}
}

func (tree *Set[T, C]) Insert(value T) {
	var node *treeNode[T]
	if tree.blackHeight == 0 {
		tree.root = &treeNode[T]{
			color: BLACK,
			value: value,
		}
		node = tree.root
	} else {
		node = tree.root
		for {
			i := tree.compare(node.value, value)
			if i < 0 {
				if node.left != nil {
					node = node.left
					continue
				}
				node.left = newNode(value, node)
				break

			} else if i > 0 {
				if node.right != nil {
					node = node.right
					continue
				}
				node.right = newNode(value, node)
				node = node.right
				break
			} else {
				return
			}
		}
	}

	tree.balanceAfterInsert(node)
}

func (tree *Set[T, C]) balanceAfterInsert(newNode *treeNode[T]) {
	if newNode.parent == nil {
		newNode.color = BLACK
		tree.blackHeight++
		return
	}
	if newNode.parent.color == BLACK {
		return
	}

	uncle := newNode.getUncle()
	if uncle.color == RED {
		newNode.parent.color = BLACK
		uncle.color = BLACK
		tree.balanceAfterInsert(newNode.parent.parent)
	} else {
		rotate(newNode)
	}
}

func rotate[T any](node *treeNode[T]) {
	parent := node.parent
	grandparent := parent.parent
	uncle := node.getUncle()

	if parent.left == node {
		if grandparent.left == parent {
			rotateLL(parent, grandparent)
		} else {
			rotateLR(node, parent, grandparent)
		}
	} else {
		if grandparent.left == parent {
			rotateRL(node, parent, grandparent, uncle)
		} else {
			rotateRR(parent, grandparent)
		}
	}
}

func rotateLL[T any](parent, grandparent *treeNode[T]) {
	grandparent.left = parent.right
	parent.parent = grandparent.parent
	grandparent.parent = parent

	//todo potencial bug (it's the same with LR)
	//pColor := parent.color
	//parent.color = grandparent.color
	//grandparent.color = pColor
	parent.color = BLACK
	grandparent.color = RED
}

func rotateLR[T any](node, parent, grandparent *treeNode[T]) {
	grandparent.left = parent.right
	parent.left = node.left
	parent.right = node.right
	node.left = parent
	node.right = grandparent

	parent.parent = node
	node.parent = grandparent.parent
	grandparent.parent = node

	node.color = BLACK
	grandparent.color = RED
}

func rotateRR[T any](parent, grandparent *treeNode[T]) {
	grandparent.right = parent.left
	parent.left = grandparent

	parent.parent = grandparent.parent
	grandparent.parent = parent

	parent.color = BLACK
	grandparent.color = RED
}

func rotateRL[T any](node, parent, grandparent, uncle *treeNode[T]) {
	temp := parent.right
	parent.right = uncle.right
	parent.left = uncle.left
	uncle.right = temp
	uncle.left = node.right
	grandparent.right = node.left
	grandparent.left = uncle
	node.left = grandparent
	node.right = parent

	node.parent = grandparent.parent
	grandparent.parent = node
	parent.parent = node

	node.color = BLACK
	grandparent.color = RED
}

// getUncle this function assumes that node has a parent that is not root
func (node *treeNode[T]) getUncle() *treeNode[T] {
	grandfather := node.parent.parent
	if grandfather.left == node.parent {
		return grandfather.right
	}
	return grandfather.left
}
