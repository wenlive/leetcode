package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) Append(data int) {
	for p.Next != nil {
		p = p.Next

	}
	var NewNode *ListNode = new(ListNode)
	NewNode.Val = data
	p.Next = NewNode
}

func NewLisNode() *ListNode {
	var head *ListNode = new(ListNode)
	for i := 0; i < 10; i++ {
		head.Append(i)
	}
	return head
}

func Cout(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head.Next.Val)
		head = head.Next
	}
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

//TODO reverse without for
