package main

import (
	"demo/leetcode"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	var nums = []int{1, -1, 0, 4}
	fmt.Println(leetcode.ThreeSum(nums))

	var a *leetcode.ListNode = leetcode.NewLisNode()
	leetcode.Cout(a)
	leetcode.ReverseBetween(a.Next, 2, 8)
	leetcode.Cout(a)
	fmt.Println("finished")

}
