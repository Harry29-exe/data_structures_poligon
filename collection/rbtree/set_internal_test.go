package rbtree

import (
	comparator "data_structures_poligon/utils"
	"testing"
)

var values = [20]int{0, -5, 5, 523, -66, -75, 7, 2, 67, 33, 55, 32, 99, 123, 223, 360, -4, 94, -88, -120}

const printDebug = false

func TestInsert(t *testing.T) {
	tree := getSetWithDefaultValues(t)

	verifyTreeStructure[int](tree, t)
}

func TestContains(t *testing.T) {
	tree := getSetWithDefaultValues(t)

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

func TestSet_Iterator(t *testing.T) {
	tree := getSetWithDefaultValues(t)
	it := tree.Iterator()
	valuesMap := make(map[int]int)
	for _, v := range values {
		valuesMap[v] = v
	}

	for it.HasNext() {
		v := it.Next()
		_, ok := valuesMap[v]
		if !ok {
			t.Error("Iterator returned value which value map does not contain")
		} else {
			delete(valuesMap, v)
		}
	}

	if len(valuesMap) != 0 {
		t.Error("Iterator does not contain all values")
	}
}

func getSetWithDefaultValues(t *testing.T) *Set[int, comparator.Compare[int]] {
	tree := NewSet[int](comparator.ComparePrimitive[int])
	for _, val := range values {
		tree.Insert(val)
		if printDebug {
			tree.print()
			print("\n\n\n")
		}
		if !tree.Contains(val) {
			t.Error("Tree does not contain inserted value:", val)
		}
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
	if nodeBlackHeight > 0 && &node.parent == nil {
		t.Error("Node with value:", node.value, "does not have parent")
	}

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
