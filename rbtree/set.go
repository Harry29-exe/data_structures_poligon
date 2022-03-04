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
				node = node.left
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
	if uncle == nil || uncle.color == BLACK {
		tree.rotate(newNode)
	} else {
		newNode.parent.color = BLACK
		newNode.parent.parent.color = RED
		uncle.color = BLACK
		tree.balanceAfterInsert(newNode.parent.parent)
	}
}

func (tree *Set[T, C]) rotate(node *treeNode[T]) {
	parent := node.parent
	grandparent := parent.parent
	uncle := node.getUncle()

	if grandparent.left == parent {
		if parent.left == node {
			tree.rotateLL(parent, grandparent)
		} else {
			tree.rotateLR(node, parent, grandparent)
		}
	} else {
		if parent.left == node {
			tree.rotateRL(node, parent, grandparent, uncle)
		} else {
			tree.rotateRR(parent, grandparent)
		}
	}
}

//todo potencial bug (it's the same with LR)
//pColor := parent.color
//parent.color = grandparent.color
//grandparent.color = pColor
func (tree *Set[T, C]) rotateLL(parent, grandparent *treeNode[T]) {
	grandparent.left = parent.right

	if grandparent.parent == nil {
		parent.parent = nil
		tree.root = parent
	} else {
		parent.parent = grandparent.parent
		if parent.parent.left == grandparent {
			parent.parent.left = parent
		} else {
			parent.parent.right = parent
		}
	}
	grandparent.parent = parent

	parent.color = BLACK
	grandparent.color = RED
}

func (tree *Set[T, C]) rotateLR(node, parent, grandparent *treeNode[T]) {
	grandparent.left = parent.left
	parent.left = node.left
	parent.right = node.right
	node.left = parent
	node.right = grandparent

	if grandparent.parent == nil {
		node.parent = nil
		tree.root = node
	} else {
		node.parent = grandparent.parent
		if node.parent.left == grandparent {
			node.parent.left = node
		} else {
			node.parent.right = node
		}
	}
	parent.parent = node

	grandparent.parent = node

	node.color = BLACK
	grandparent.color = RED
}

func (tree *Set[T, C]) rotateRR(parent, grandparent *treeNode[T]) {
	grandparent.right = parent.left
	parent.left = grandparent

	if grandparent == nil {
		parent.parent = nil
		tree.root = parent
	} else {
		parent.parent = grandparent.parent
		if parent.parent.left == grandparent {
			parent.parent.left = parent
		} else {
			parent.parent.right = parent
		}
	}
	grandparent.parent = parent

	parent.color = BLACK
	grandparent.color = RED
}

func (tree *Set[T, C]) rotateRL(node, parent, grandparent, uncle *treeNode[T]) {
	if uncle != nil {
		temp := parent.right
		parent.right = uncle.right
		parent.left = uncle.left
		uncle.right = temp
		uncle.left = node.right
		grandparent.left = uncle
	}
	grandparent.right = node.left
	node.left = grandparent
	node.right = parent

	if grandparent.parent == nil {
		node.parent = nil
		tree.root = node
	} else {
		node.parent = grandparent.parent
		if node.parent.left == grandparent {
			node.parent.left = node
		} else {
			node.parent.right = node
		}
	}
	grandparent.parent = node
	parent.parent = node

	node.color = BLACK
	grandparent.color = RED
}

func (tree *Set[T, C]) print() {
	if tree.root != nil {
		printNode(tree.root, 0)
	}
}

func printNode[T any](node *treeNode[T], depth uint) {
	if node.right != nil {
		printNode(node.right, depth+1)
	}
	spacing := ""
	for i := uint(0); i < depth; i++ {
		spacing = spacing + "    "
	}
	print(spacing)
	var c string
	if node.color == RED {
		c = "R"
	} else {
		c = "B"
	}
	print(node.value)
	println(c)
	if node.left != nil {
		printNode(node.left, depth+1)
	}
}

// getUncle this function assumes that node has a parent that is not root
func (node *treeNode[T]) getUncle() *treeNode[T] {
	grandfather := node.parent.parent
	if grandfather.left == node.parent {
		return grandfather.right
	}
	return grandfather.left
}
