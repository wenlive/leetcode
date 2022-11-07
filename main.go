package main

import (
	"demo/leetcode"
	"fmt"
)

func TestreverseLinkedList() {
	// var a *leetcode.ListNode = leetcode.NewLisNode()
	// leetcode.Cout(a)

	b := leetcode.Reverse(leetcode.NewLisNode())
	leetcode.Cout(b)

	// b2 := leetcode.ReverseWithoutLoop(leetcode.NewLisNode())
	// leetcode.Cout(b2)

	// c := leetcode.ReverseBetween(leetcode.NewLisNode(), 2, 8)
	// leetcode.Cout(c)
	// c2 := leetcode.ReverseBetweenV2(leetcode.NewLisNode(), 2, 8)
	// leetcode.Cout(c2)
	// d := leetcode.ReverseBetweenWithoutLoop(leetcode.NewLisNode(), 2, 8)
	// leetcode.Cout(d)

	// e := leetcode.ReverseN(leetcode.NewLisNode(), 4)
	// leetcode.Cout(e)

	// fmt.Println("finished")

}

func Test3Sum() {
	fmt.Println("hello world")
	// var nums = []int{1, -1, 0, 4}
	// fmt.Println(leetcode.ThreeSum(nums))
}

func TestBinaryTree() {
	root := leetcode.NewMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	res := leetcode.LevelOrder(root)
	for _, i := range res {
		fmt.Println(i)
	}
}

func Testpalindrome() {
	// var nums = []int{-1, 1, 1, 4}
	// fmt.Println(leetcode.RemoveDuplicates(nums))
	// fmt.Println(nums)

	a := leetcode.NewLisNodeWith([]int{1, 2, 3, 2, 1})
	leetcode.Cout(a)
	// fmt.Println(leetcode.IsPalindrome(a))

	fmt.Println(leetcode.IsPalindromeV2(a))
}

func main() {
	// TestreverseLinkedList()

	// Testpalindrome()

	// fmt.Println(leetcode.MaxCharInString("aabbbbcccaabbbbccc"))

	TestBinaryTree()
}
