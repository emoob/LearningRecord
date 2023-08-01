package list_node

import "testing"

func TestSingleListNode(T *testing.T) {
	n := NewNode()
	for i := 1; i < 11; i++ {
		n.Insertlist(i, i)
	}
	n.DelNode(1)
	n.DelNode(1)
	n.DelNode(1)
	n.Print()
}
