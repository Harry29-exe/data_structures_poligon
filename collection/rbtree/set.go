package rbtree

import (
	"data_structures_poligon/collection"
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
			if i > 0 {
				if node.left != nil {
					node = node.left
					continue
				}
				node.left = newNode(value, node)
				node = node.left
				break

			} else if i < 0 {
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

func (tree *Set[T, C]) Contains(value T) bool {
	if tree.root == nil {
		return false
	}
	node := tree.root
	for {
		v := tree.compare(node.value, value)
		if v < 0 {
			if node.right != nil {
				node = node.right
				continue
			} else {
				return false
			}
		} else if v > 0 {
			if node.left != nil {
				node = node.left
				continue
			} else {
				return false
			}
		} else {
			return true
		}
	}
}

func (tree *Set[T, C]) Iterator() collection.Iterator[T] {
	iterator := &setIterator[T]{
		next:      tree.root,
		lastEvent: DOWN,
	}

	return iterator
}

func (tree *Set[T, C]) Delete(t T) {
	//TODO implement me
	panic("implement me")
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

func (tree *Set[T, C]) rotateLL(parent, grandparent *treeNode[T]) {
	grandparent.left = parent.right
	if grandparent.left != nil {
		grandparent.left.parent = grandparent
	}

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
	parent.right = grandparent
	grandparent.parent = parent

	parent.color = BLACK
	grandparent.color = RED
}

func (tree *Set[T, C]) rotateLR(node, parent, grandparent *treeNode[T]) {
	grandparent.left = node.right
	if grandparent.left != nil {
		grandparent.left.parent = grandparent
	}
	parent.right = node.left
	if parent.right != nil {
		parent.right.parent = parent
	}
	node.left = parent
	parent.parent = node

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
	node.right = grandparent
	grandparent.parent = node

	node.color = BLACK
	grandparent.color = RED
}

//func (tree *Set[T, C]) rotateLR(node, parent, grandparent *treeNode[T]) {
//	grandparent.left = parent.left
//	if grandparent.left != nil {
//		grandparent.left.parent = grandparent
//	}
//	parent.left = node.left
//	if parent.left != nil {
//		parent.left.parent = parent
//	}
//	parent.right = node.right
//	if parent.right != nil {
//		parent.right.parent = parent
//	}
//	node.left = parent
//	parent.parent = node
//
//	if grandparent.parent == nil {
//		node.parent = nil
//		tree.root = node
//	} else {
//		node.parent = grandparent.parent
//		if node.parent.left == grandparent {
//			node.parent.left = node
//		} else {
//			node.parent.right = node
//		}
//	}
//	node.right = grandparent
//	grandparent.parent = node
//
//	node.color = BLACK
//	grandparent.color = RED
//}

func (tree *Set[T, C]) rotateRR(parent, grandparent *treeNode[T]) {
	grandparent.right = parent.left
	if grandparent.right != nil {
		grandparent.right.parent = grandparent
	}

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
	parent.left = grandparent
	grandparent.parent = parent

	parent.color = BLACK
	grandparent.color = RED
}

func (tree *Set[T, C]) rotateRL(node, parent, grandparent, uncle *treeNode[T]) {
	grandparent.right = node.left
	if grandparent.right != nil {
		grandparent.right.parent = grandparent
	}
	parent.left = node.right
	if parent.left != nil {
		parent.left.parent = parent
	}

	node.right = parent
	parent.parent = node

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
	node.left = grandparent
	grandparent.parent = node

	node.color = BLACK
	grandparent.color = RED
}

//func (tree *Set[T, C]) rotateRL(node, parent, grandparent, uncle *treeNode[T]) {
//	if uncle != nil {
//		temp := parent.right
//		parent.right = uncle.right
//		if parent.right != nil {
//			parent.right.parent = parent
//		}
//		parent.left = uncle.left
//		if parent.left != nil {
//			parent.left.parent = parent
//		}
//		uncle.right = temp
//		if uncle.right != nil {
//			uncle.right.parent = uncle
//		}
//		uncle.left = node.right
//		if uncle.left != nil {
//			uncle.left.parent = uncle
//		}
//	} else {
//		parent.left = nil
//		parent.right = nil
//	}
//
//	grandparent.right = node.left
//	if grandparent.right != nil {
//		grandparent.right.parent = grandparent
//	}
//
//	node.right = parent
//	parent.parent = node
//
//	if grandparent.parent == nil {
//		node.parent = nil
//		tree.root = node
//	} else {
//		node.parent = grandparent.parent
//		if node.parent.left == grandparent {
//			node.parent.left = node
//		} else {
//			node.parent.right = node
//		}
//	}
//	node.left = grandparent
//	grandparent.parent = node
//
//	node.color = BLACK
//	grandparent.color = RED
//}

func (tree *Set[T, C]) print() {
	if tree.root != nil {
		printNode(tree.root, 0)
	}
}

func newNode[T any](value T, parent *treeNode[T]) *treeNode[T] {
	return &treeNode[T]{value: value, color: RED, parent: parent}
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

// getUncle this function assumes that next has a parent that is not root
func (node *treeNode[T]) getUncle() *treeNode[T] {
	grandfather := node.parent.parent
	if grandfather.left == node.parent {
		return grandfather.right
	}
	return grandfather.left
}

type lastIterEvent = uint8

const (
	DOWN lastIterEvent = iota
	LEFT_UP
	RIGHT_UP
)

type setIterator[T any] struct {
	next      *treeNode[T]
	lastEvent lastIterEvent
}

func (s *setIterator[T]) HasNext() bool {
	return s.next != nil
}

func (s *setIterator[T]) Next() T {
	if s.next == nil {
		panic("Iterator next is nil")
	}
	nodeToReturn := s.next

	temp := nodeToReturn
	for temp != nil {
		switch s.lastEvent {
		case DOWN:
			temp = s.lastEvDown(temp)
		case LEFT_UP:
			temp = s.lastEvLeftUp(temp)
		case RIGHT_UP:
			temp = s.lastEvRightUp(temp)
		}
	}

	return nodeToReturn.value
}

func (s *setIterator[T]) lastEvDown(node *treeNode[T]) *treeNode[T] {
	if node.left != nil {
		s.lastEvent = DOWN
		s.next = node.left
	} else if node.right != nil {
		s.lastEvent = DOWN
		s.next = node.right
	} else if node.parent.left == node {
		s.lastEvent = LEFT_UP
		return node.parent
	} else {
		s.lastEvent = RIGHT_UP
		return node.parent
	}

	return nil
}

func (s *setIterator[T]) lastEvLeftUp(node *treeNode[T]) *treeNode[T] {
	if node.right != nil {
		s.lastEvent = DOWN
		s.next = node.right
	} else if node.parent == nil {
		s.next = nil
	} else if node.parent.left == node {
		s.lastEvent = LEFT_UP
		return node.parent
	} else {
		s.lastEvent = RIGHT_UP
		return node.parent
	}

	return nil
}

func (s *setIterator[T]) lastEvRightUp(node *treeNode[T]) *treeNode[T] {
	if node.parent == nil {
		s.next = nil
	} else if node.parent.left == node {
		s.lastEvent = LEFT_UP
		return node.parent
	} else {
		s.lastEvent = RIGHT_UP
		return node.parent
	}

	return nil
}
