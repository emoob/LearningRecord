package sequence

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	c, err := NewList([]int{1, 999, 555, 88})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Sort()
	fmt.Printf("排序后数组：%d \n", c.E)
	if index, ok := c.GettingElements(999); ok {
		fmt.Printf("所在下标：%d \n", index)
	}

}
