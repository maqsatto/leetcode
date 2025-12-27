package leet2

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	extra := 0
	for l1 != nil || l2 != nil || extra != 0 {
		sum := 0
		sum += extra
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		extra = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next

	}
	return head.Next
}

func TestFunc() {
	// l1 = [2,4,3]
	l1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	// l2 = [5,6,4]
	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	fmt.Println(addTwoNumbers(l1, l2), addTwoNumbers(l1, l2).Next, addTwoNumbers(l1, l2).Next.Next)
}
