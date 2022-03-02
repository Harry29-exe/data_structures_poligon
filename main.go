package main

import "data_structures_poligon/linked_list"

func main() {
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

func test(l linked_list.LinkedList[int]) {
	l.Size = 42
	println(l.Size)
}
