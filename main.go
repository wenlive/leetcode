package main

import (
	"demo/leetcode"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	var nums = []int{1, -1, 0, 4}
	fmt.Println(leetcode.ThreeSum(nums))

	// var a *leetcode.ListNode = leetcode.NewLisNode()
	// leetcode.Cout(a)
	b := leetcode.Reverse(leetcode.NewLisNode())
	leetcode.Cout(b)

	b2 := leetcode.ReverseWithoutLoop(leetcode.NewLisNode())
	leetcode.Cout(b2)

	// c := leetcode.ReverseBetween(leetcode.NewLisNode(), 2, 8)
	// leetcode.Cout(c)
	// d := leetcode.ReverseBetweenV2(leetcode.NewLisNode(), 2, 8)
	// leetcode.Cout(d)

	// fmt.Println("finished")

}
