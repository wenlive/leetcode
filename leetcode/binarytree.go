package leetcode

import "math"

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func NewMaximumBinaryTree(nums []int) *BinaryTree {

	if len(nums) == 0 {
		return nil
	}
	var max int = int(math.MinInt64)
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	root := BinaryTree{Val: max}
	if index < 0 {
		root.Left = NewMaximumBinaryTree(nums[:0])
	} else {
		root.Left = NewMaximumBinaryTree(nums[0:index])
	}

	// root.Right = NewBinaryTree(nums[index+1 : len(nums)-1])
	if index+1 > len(nums) {
		root.Right = NewMaximumBinaryTree(nums[:0])
	} else {
		root.Right = NewMaximumBinaryTree(nums[index+1:])
	}

	return &root
}

func LevelOrder(root *BinaryTree) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*BinaryTree{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*BinaryTree{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
