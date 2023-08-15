package stack

import (
	"log"
)

type ArrayStack struct {
	E        []int
	Capacity int
	Top      int
}

func NewArrayStack(cap int) *ArrayStack {
	return &ArrayStack{
		E:        nil,
		Capacity: cap,
		Top:      -1,
	}
}

// Push 入栈
func (s *ArrayStack) Push(e int) int {
	if len(s.E) >= s.Capacity {
		log.Print("容量超过初始大小", s.Capacity)
		return 0
	}
	s.Top = e
	s.E = append([]int{e}, s.E...)
	return 1
}

// Out 出栈
func (s *ArrayStack) Out(e int) int {
	if len(s.E) <= 0 {
		log.Print("当前栈内没有元素")
		return 0
	}
	for i := 0; i < len(s.E); i++ {
		if e == s.E[i] {
			s.Top = s.E[i+1]
			s.E = append([]int{s.E[i+1]}, s.E[i+2:]...)
			return 1
		}
	}
	return 0
}
