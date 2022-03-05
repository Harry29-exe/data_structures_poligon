package rbtree

import (
	comparator "data_structures_poligon/utils"
	"testing"
)

var values = [20]int{0, -5, 5, 523, -66, -75, 7, 2, 67, 33, 55, 32, 99, 123, 223, 360, -4, 94, -88, -120}

func TestInsert(t *testing.T) {
	tree := getSetWithDefaultValues()

	verifyTreeStructure[int](tree, t)
}

func TestContains(t *testing.T) {
	tree := getSetWithDefaultValues()

	for _, v := range values {
		if !tree.Contains(v) {
			t.Error("Tree should contain:", v, "but tree.Contains return false")
		}
	}

	if tree.Contains(10000) {
		t.Error()
	}
	if tree.Contains(-10000) {
		t.Error()
	}
}

func getSetWithDefaultValues() *Set[int, comparator.Compare[int]] {
	tree := NewSet[int](comparator.ComparePrimitive[int])
	for _, val := range values {
		tree.Insert(val)
	}
	return tree
}

func verifyTreeStructure[T any, C comparator.Compare[T]](tree *Set[T, C], t *testing.T) {
	if tree.root == nil {
		return
	}
	if tree.root.color == RED {
		t.Error("Tree root is RED")
	}

	verifyNode(tree.root, 0, t)
}

func verifyNode[T any](node *treeNode[T], nodeBlackHeight uint, t *testing.T) uint {
	if node.color == RED {
		if node.left != nil && node.left.color == RED {
			t.Error("Node with value ", node.value, " is RED and has left RED child")
		} else if node.right != nil && node.right.color == RED {
			t.Error("Node with value ", node.value, " is RED and has right RED child")
		}
	}

	if node.color == BLACK {
		nodeBlackHeight++
	}

	var leftH, rightH uint
	if node.left != nil {
		leftH = verifyNode(node.left, nodeBlackHeight, t)
	} else {
		leftH = nodeBlackHeight
	}

	if node.right != nil {
		rightH = verifyNode(node.right, nodeBlackHeight, t)
	} else {
		rightH = nodeBlackHeight
	}

	if leftH != rightH {
		t.Error("Path to nil nodes have different black heights")
	}

	return leftH
}
