package list_node

import "fmt"

type ListNode struct {
	Structure int
	Node      *ListNode
}

func (n *ListNode) NewNode(node ListNode) *ListNode {
	return &node
}

// Insertlist 链表的插入
func (n *ListNode) Insertlist(head *ListNode, e int, index int) int {
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

// PrintList 查看所有节点
func (n *ListNode) PrintList() {
	head := n.Node
	for head != nil {
		fmt.Printf("%v\n", head.Structure)
		head = head.Node
	}
}
func (n *ListNode) DelNode(index int) int {
	if index < 1 {
		return 0
	}
	head := n.Node
	if index == 1 {
		// 删除头节点
		n.Node = n.Node.Node
		return 1
	}
	// 遍历到index - 1所在的节点
	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Node
	}
	// 如果current=nil则超出链表长度
	if current == nil || current.Node == nil {
		// 节点不存在
		return 0
	}
	// 删除节点
	current.Node = current.Node.Node
	return 1
}
