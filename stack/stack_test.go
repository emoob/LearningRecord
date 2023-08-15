package stack

import (
	"log"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewArrayStack(10)
	log.Printf("容量:%d\n", stack.Capacity)
	log.Printf("顶部元素:%d\n", stack.Top)
	log.Printf("栈内元素:%d\n", stack.E)
	for i := 0; i < 11; i++ {
		stack.Push(i)
	}
	log.Printf("顶部元素:%d\n", stack.Top)
	log.Printf("栈内元素:%d\n", stack.E)
	stack.Out(7)
	log.Printf("栈内元素:%d\n", stack.E)
	log.Printf("顶部元素:%d\n", stack.Top)
}
