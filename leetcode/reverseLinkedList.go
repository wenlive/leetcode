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
	head.Val = 1
	for i := 2; i < 10; i++ {
		head.Append(i)
	}
	return head
}

func Cout(head *ListNode) {
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println(head.Val)
	fmt.Println("----------")
}

func Reverse(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

func ReverseWithoutLoop(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//利用递归将head不停前移
	last := ReverseWithoutLoop(head.Next)
	//将后面的next指向当前位置
	head.Next.Next = head
	//清空原有正向连接
	head.Next = nil
	return last
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

func ReverseBetweenV2(head *ListNode, m int, n int) *ListNode {
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
	var behind *ListNode
	for i := 0; i < n-m+1; i++ {
		next := cur.Next
		// if i == n-m {
		// 	if next == nil {
		// 		fmt.Println("touch the nil bound")
		// 	}
		// }
		cur.Next = behind
		behind = cur
		cur = next
	}
	pre.Next.Next = cur
	pre.Next = behind
	return newHead.Next
}
