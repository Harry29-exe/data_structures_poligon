package main

import (
	"data_structures_poligon/linked_list"
	"data_structures_poligon/rbtree"
	comparator "data_structures_poligon/utils"
	"time"
)

func main() {
	structPointerTest()
	rbTreeSetTest()
}

type testStruct struct {
	name string
	age  uint8
}

func structPointerTest() {
	a1 := testStruct{name: "abcde", age: 1}
	a1p := &a1
	a2 := testStruct{name: "abcde", age: 1}
	a3 := testStruct{name: "abcde", age: 2}

	a1p2 := &a1
	time0 := time.Now()
	a1_a2 := a1 == a2
	a1_a2 = a1 == a2
	a1_a2 = a1 == a2
	a1_a2 = a1 == a2
	a1_a2 = a1 == a2

	time_a1_a2 := time.Now()
	a1_a3 := a1 == a3
	a1_a3 = a1 == a3
	a1_a3 = a1 == a3
	a1_a3 = a1 == a3
	a1_a3 = a1 == a3

	time_a1_a3 := time.Now()
	a1_a1p := a1p2 == a1p
	a1_a1p = a1p2 == a1p
	a1_a1p = a1p2 == a1p
	a1_a1p = a1p2 == a1p
	a1_a1p = a1p2 == a1p
	time_a1_a1p := time.Now()

	println(time_a1_a2.UnixNano()-time0.UnixNano(),
		time_a1_a3.UnixNano()-time_a1_a2.UnixNano(),
		time_a1_a1p.UnixNano()-time_a1_a3.UnixNano())
	println(a1_a2 && a1_a3 && a1_a1p)
}

func linkedListTest() {
	list := linked_list.New[int]()

	zero, err := list.Get(0)
	if err != nil {
		println(zero)
	}

	list.Insert(15)
	list.Insert(2)

	for i := uint(0); i < list.Size; i++ {
		item, _ := list.Get(i)
		println(item)
	}

	test(list)
	println(list.Size)

	var i int = -5325
	println(int8(i))
}

func rbTreeSetTest() {
	tree := rbtree.NewSet[int](comparator.ComparePrimitive[int])
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(7)

}

func test(l linked_list.LinkedList[int]) {
	l.Size = 42
	println(l.Size)
}
