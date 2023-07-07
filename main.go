package main

import (
	"fmt"
	"structure/sequence"
)

// @sum 高斯求合
func sum(a, n int) int {
	return (a + n) * n / 2
}

//	func main() {
//		b := binary.Binary{
//			Nums:     []int{0, 2, 5, 6, 9, 15, 156, 555, 985, 101},
//			NumsSize: 10,
//			Target:   5,
//		}
//		fmt.Println(b.Search())
//	}

func main() {
	d := &sequence.List{}
	c, err := d.NewList([]int{1, 999, 555, 88})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Sort()
	fmt.Print(c.E)
	if index, ok := c.GettingElements(999); ok {
		fmt.Print(index)
	}

}
