package leet21

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	cur := mergedList
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return mergedList.Next
}
func printList(head *ListNode) {
	if head == nil {
		fmt.Println("nil")
	}
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

func TestMergeTwoLists() {
	// list1 = [1,2,4]
	list1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	list1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	list1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	// list2 = [1,3,4]
	list2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	list2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	list2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	printList(mergeTwoLists(list1, list2))
}
