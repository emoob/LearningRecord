package stack

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewArrayStack(10)
	log.Printf("容量:%d 顶部元素:%d 栈内元素:%d 栈内元素个数:%d",
		stack.Capacity,
		stack.Top,
		stack.E,
		len(stack.E))
	for i := 2; i < 12; i++ {
		stack.Push(i)
	}
	for i := 0; i < 11; i++ {
		stack.Out()
	}
	log.Printf("容量:%d 顶部元素:%d 栈内元素:%d 栈内元素个数:%d",
		stack.Capacity,
		stack.Top,
		stack.E,
		len(stack.E))

	data := []int{0, 2, 2, 3, 4, 4}
	for i := 0; i < len(data); i++ {
		if data[i] == data[i+1] {

		}
	}
}
