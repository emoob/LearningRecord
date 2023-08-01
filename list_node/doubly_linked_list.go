package list_node

import "fmt"

type DoublyNode struct {
	E    int
	next *DoublyNode
	prev *DoublyNode
}

func NewDoublyNode() *DoublyNode {
	return &DoublyNode{
		E:    0,
		next: nil,
		prev: nil,
	}
}
func (d *DoublyNode) Insertlist(e int, index int) int {
	if index < 1 {
		return 0
	}
	current := d
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}
	if current == nil {
		return 0
	}
	// 在当前节点存在的情况下，下一个节点就是当前节点的下一个节点
	// 上一个节点就是当前节点
	newNode := &DoublyNode{
		E:    e,
		next: current.next,
		prev: current,
	}
	// 将新节点插入到链表中
	current.next = newNode
	// 如果新节点的下一个节点不是 nil，
	// 则将新节点的下一个节点的上一个节点指针指向新节点
	if newNode.next != nil {
		newNode.next.prev = newNode
	}
	return 1
}

// Print 查看所有节点数据
func (d *DoublyNode) Print() {
	current := d.next
	for current != nil {
		fmt.Printf("当前节点数据%v \n", current.E)
		current = current.next
	}
}
