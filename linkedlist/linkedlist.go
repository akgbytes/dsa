package linkedlist

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len  int
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.len {
		return -1
	}

	current := ll.head
	for _ = range index {
		current = current.Next
	}

	return current.Val
}

func (ll *LinkedList) AddAtHead(val int) {
	n := &ListNode{Val: val}
	n.Next = ll.head
	ll.head = n
	ll.len++
}

func (ll *LinkedList) AddAtTail(val int) {
	n := &ListNode{Val: val}

	if ll.head == nil {
		ll.head = n
		ll.len++
		return
	}

	current := ll.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = n
	ll.len++
}

func (ll *LinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > ll.len {
		return
	}

	if index == 0 {
		ll.AddAtHead(val)
		return
	}

	current := ll.head
	for _ = range index - 1 {
		current = current.Next
	}

	n := &ListNode{Val: val}
	n.Next = current.Next
	current.Next = n
	ll.len++
}

func (ll *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= ll.len {
		return
	}

	sentinel := &ListNode{Next: ll.head}
	current := sentinel

	for _ = range index {
		current = current.Next
	}

	current.Next = current.Next.Next

	ll.head = sentinel.Next
	ll.len--
}

func (ll *LinkedList) Print() string {
	res := ""
	current := ll.head

	if current == nil {
		return res + "empty"
	}

	res += strconv.Itoa(current.Val)

	for current.Next != nil {
		current = current.Next
		res = res + " -> " + strconv.Itoa(current.Val)
	}

	return res
}

func (ll *LinkedList) GetHead() *ListNode {
	return ll.head
}
