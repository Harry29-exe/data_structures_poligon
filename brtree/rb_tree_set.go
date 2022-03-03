package brtree

import . "data_structures_poligon/utils"

type leafColor = uint8

const (
	RED leafColor = iota
	BLACK
)

type RBTreeSet[T any, C Compare[T]] struct {
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

func newNode[T](value T, parent *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{value: value, color: RED, parent: parent}
}

func NewPrimitiveSet[T any]() *RBTreeSet[T] {
	return &RBTreeSet[T]{
		root:        nil,
		blackHeight: uint(0),
	}
}

func (tree *RBTreeSet[T]) Insert(value T) {
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

func (tree RBTreeSet[T]) balanceAfterInsert(newNode *treeNode[T]) {
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

	}
}

func (tree RBTreeSet[T]) rotate(node *treeNode[T]) {
	parent := node.parent
	grandparent := parent.parent

	if parent.left == node {
		if grandparent.left == parent {

		}
	}
}

func rotateLL[T](node, parent, grandparent *treeNode[T]) {
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

func rotateLR[T](node, parent, grandparent *treeNode[T]) {
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

func rotateRR[T](node, parent, grandparent *treeNode[T]) {
	grandparent.right = parent.left
	parent.left = grandparent

	parent.parent = grandparent.parent

}

// getUncle this function assumes that node has a parent that is not root
func (node *treeNode[T]) getUncle() *treeNode[T] {
	grandfather := node.parent.parent
	if grandfather.left == node.parent {
		return grandfather.right
	}
	return grandfather.left
}
