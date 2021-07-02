package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(value int) *ListNode {
	if value < 0 {
		return nil
	}
	if value < 10 {
		return &ListNode{Val: value}
	}
	return &ListNode{Val: value % 10, Next: New(value / 10)}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// check for nil
	if l1 == nil {
		// it's ok if l2 is nil here, because it is expected it can be nil at this point.
		return l2
	}
	if l2 == nil {
		return l1
	}

	if (l1.Val + l2.Val) > 9 {
		if( l1.Next == nil){
			l1.Next = &ListNode{Val: 1}
		} else {
			l1.Next.Val += 1
		}
		if l2.Next == nil {
			l2.Next = &ListNode{}
		}
	}

	return &ListNode{
		Val:  (l1.Val + l2.Val) % 10,
		Next: addTwoNumbers(l1.Next, l2.Next),
	}
}
