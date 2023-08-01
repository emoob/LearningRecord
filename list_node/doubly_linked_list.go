package list_node

type DoublyNode struct {
	E    int
	next *DoublyNode
	prev *DoublyNode
}

func NewDoublyNode(h DoublyNode) {
	h.next = nil
	h.prev = nil
}
