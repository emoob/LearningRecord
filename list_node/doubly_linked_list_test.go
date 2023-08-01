package list_node

import "testing"

func TestDoublyNode(t *testing.T) {
	d := NewDoublyNode()
	d.Insertlist(10, 1)
	d.Insertlist(11, 1)
	d.Insertlist(12, 1)
	d.Print()
}
