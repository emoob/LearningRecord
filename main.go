package main

import (
	"structure/list_node"
)

//	func main() {
//		b := binary.Binary{
//			Nums:     []int{0, 2, 5, 6, 9, 15, 156, 555, 985, 101},
//			NumsSize: 10,
//			Target:   5,
//		}
//		fmt.Println(b.Search())
//	}

//func main() {
//	d := &sequence.List{}
//	c, err := d.NewList([]int{1, 999, 555, 88})
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	c.Sort()
//	fmt.Print(c.E)
//	if index, ok := c.GettingElements(999); ok {
//		fmt.Print(index)
//	}
//
//}

func main() {
	node := &list_node.ListNode{
		Structure: 0,
		Node:      nil,
	}
	for i := 1; i < 11; i++ {
		node.Insertlist(node, i, 1)
	}
	node.DelNode(5)
	node.PrintList()

}
