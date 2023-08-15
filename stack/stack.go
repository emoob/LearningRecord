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
		E:        []int{},
		Capacity: cap,
		Top:      -1,
	}
}

// Push 入栈
func (s *ArrayStack) Push(e int) bool {
	if len(s.E) == s.Capacity {
		log.Print("容量超过初始大小", s.Capacity)
		return false
	}
	s.Top = e
	s.E = append([]int{e}, s.E...)
	return true
}

// Out 出栈
func (s *ArrayStack) Out() bool {
	if len(s.E) == 0 {
		log.Print("当前栈内没有元素")
		return false
	}
	if len(s.E) == 1 {
		s.Top = -1
		s.E = nil
		return true
	}
	s.Top = s.E[1]
	s.E = append([]int{s.E[1]}, s.E[2:]...)
	return true
}
