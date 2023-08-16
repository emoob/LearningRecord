package stack

import "log"

type Queue struct {
	E     []int
	Cap   int
	front int
	rear  int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		E:     []int{},
		Cap:   cap,
		rear:  0,
		front: 0,
	}
}

func (q *Queue) EeQueue(e int) bool {
	if len(q.E) >= q.Cap {
		log.Printf("无法入队,当前对内个数%d", len(q.E))
		return false
	}
	q.rear = (q.rear + 1) % q.Cap
	q.E = append(q.E, e)
	return true
}
func (q *Queue) DeQueue() bool {
	if len(q.E) < 1 {
		log.Print("当前队内没有数据")
		return false
	}
	q.E = q.E[1:]
	q.front = (q.front + 1) % q.Cap
	return true
}

func (q *Queue) Print() {
	log.Printf("当前元素%d rear：%d front:%d", q.E, q.rear, q.front)
}
