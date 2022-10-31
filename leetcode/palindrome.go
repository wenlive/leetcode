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
