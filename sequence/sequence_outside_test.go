package sequence

import (
	"log"
	"testing"
)

func TestList(t *testing.T) {
	c, err := NewList([]int{1, 999, 555, 88})
	if err != nil {
		log.Printf(err.Error())
		return
	}
	c.Sort()
	log.Printf("排序后数组：%d \n", c.E)
	if index, ok := c.GettingElements(999); ok {
		log.Printf("所在下标：%d \n", index)
	}

}
