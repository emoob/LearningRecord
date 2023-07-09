package list_node

type E int

type ListNode struct {
	Structure E
	Node      *ListNode
}

func (n *ListNode) NewNode(node ListNode) *ListNode {
	return &node
}

// Insertlist 链表的插入
func (n *ListNode) Insertlist(head *ListNode, e E, index int) int {
	if index < 1 {
		return 0
	}
	current := head
	// 遍历到index所在的节点
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Node
	}
	// 如果current=nil则超出链表长度
	if current == nil {
		return 0
	}
	// 创建一个新的结构，并把数据插入
	newNode := &ListNode{Structure: e, Node: current.Node}
	current.Node = newNode
	return 1
}
