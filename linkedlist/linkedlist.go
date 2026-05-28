package linkedlist

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len  int
}

func (ll *LinkedList) Insert(val int) {
	ListNode := &ListNode{Val: val}
	if ll.head == nil {
		ll.head = ListNode
		ll.len++
		return
	}

	curr := ll.head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = ListNode
	ll.len++
}

func (ll *LinkedList) InsertAt(val int, idx int) {
	if ll.len < idx || idx < 0 {
		return
	}

	ListNode := &ListNode{Val: val}

	if idx == 0 {
		ListNode.Next = ll.head
		ll.head = ListNode
		ll.len++
		return
	}

	curr := ll.head

	for _ = range idx - 1 {
		curr = curr.Next
	}

	ListNode.Next = curr.Next
	curr.Next = ListNode
	ll.len++

}

func (ll *LinkedList) Delete(val int) {
	if ll.head == nil {
		return
	}

	sentinel := &ListNode{}
	sentinel.Next = ll.head

	curr := sentinel
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			ll.head = sentinel.Next
			ll.len--
			return
		}
		curr = curr.Next
	}
}

func (ll *LinkedList) DeleteAt(idx int) {
	if idx >= ll.len || idx < 0 {
		return
	}

	if idx == 0 {
		ll.head = ll.head.Next
		ll.len--
		return
	}

	curr := ll.head

	for _ = range idx - 1 {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
	ll.len--
}

func (ll *LinkedList) Find(val int) int {
	idx := 0
	for curr := ll.head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			return idx
		}
		idx++
	}
	return -1
}

func (ll *LinkedList) Print() {
	for curr := ll.head; curr != nil; curr = curr.Next {
		fmt.Printf("%d -> ", curr.Val)
	}
	fmt.Print("nil\n")
}
