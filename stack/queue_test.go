package stack

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue(10)
	for i := 0; i < 15; i++ {
		q.EeQueue(i)
	}
	q.Print()
	q.DeQueue()
	q.Print()
}
