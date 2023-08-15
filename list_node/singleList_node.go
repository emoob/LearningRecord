package list_node

import "fmt"

type SingleListNode struct {
	Structure int
	Node      *SingleListNode
}

func NewNode() *SingleListNode {
	return &SingleListNode{
		Structure: 0,
		Node:      nil,
	}
}

// Insertlist 插入
func (n *SingleListNode) Insertlist(e int, index int) bool {
	if index < 1 {
		return false
	}
	current := n
	// 遍历到index所在的节点
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Node
	}
	// 如果current=nil则超出链表长度
	if current == nil {
		return false
	}
	// 创建一个新的结构，并把数据插入
	newNode := &SingleListNode{Structure: e, Node: current.Node}
	current.Node = newNode
	return true
}

// Print 查看所有节点数据
func (n *SingleListNode) Print() {
	head := n.Node
	for head != nil {
		fmt.Printf("当前节点数据:%v\n", head.Structure)
		head = head.Node
	}
}

// DelNode 删除某个节点
func (n *SingleListNode) DelNode(index int) bool {
	if index < 1 {
		return false
	}
	head := n.Node
	if index == 1 {
		n.Node = n.Node.Node
		return true
	}
	current := head
	// 遍历到这个节点所在得位置
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Node
	}
	// 节点超出链表长度
	if current == nil || current.Node == nil {
		return false
	}
	// 删除节点
	current.Node = current.Node.Node
	return true
}
