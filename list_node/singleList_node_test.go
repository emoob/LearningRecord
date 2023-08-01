package list_node

import "testing"

func TestSingleListNode(T *testing.T) {
	node := &SingleListNode{
		Structure: 0,
		Node:      nil,
	}
	for i := 1; i < 11; i++ {
		node.Insertlist(node, i, 1)
	}
	node.DelNode(5)
	node.PrintList()
}
