package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// l1 和 l2 为当前遍历的节点，carry 为进位
func pushTwo(l1, l2 *ListNode, carry int) *ListNode {
	// 递归边界：l1 和 l2 都是空节点
	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &ListNode{Val: carry} // 如果进位了，就额外创建一个节点
		}
		return nil
	}
	// 如果 l1 是空的，那么此时 l2 一定不是空节点
	if l1 == nil {
		// 交换 l1 与 l2，保证 l1 非空，从而简化代码
		l1, l2 = l2, l1
	}
	// 节点值和进位加在一起
	carry += l1.Val
	if l2 != nil {
		// 节点值和进位加在一起
		carry += l2.Val
		// 下一个节点
		l2 = l2.Next
	}
	// 每个节点保存一个数位
	l1.Val = carry % 10
	// 进位
	l1.Next = pushTwo(l1.Next, l2, carry/10)
	return l1
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	return pushTwo(l1, l2, 0)
}
