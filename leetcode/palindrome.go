package leetcode

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func NewLisNodeWith(nums []int) *ListNode {
	var pre *ListNode = new(ListNode)
	for _, v := range nums {
		pre.Append(v)
	}
	return pre.Next
}

var left *ListNode

func IsPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)

}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	ans := res && (right.Val == left.Val)
	left = left.Next
	return ans
}

func IsPalindromeV2(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := ReverseToNewLinkedList(head)
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func ReverseToNewLinkedList(head *ListNode) *ListNode {
	var last *ListNode
	for head != nil {
		temp := ListNode{Val: head.Val}
		head = head.Next
		temp.Next = last
		last = &temp
	}

	return last
}
