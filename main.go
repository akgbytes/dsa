package main

import "github.com/akgbytes/dsa/linkedlist"

func main() {

	ll := linkedlist.LinkedList{}

	ll.AddAtHead(5)
	ll.AddAtHead(4)
	ll.AddAtHead(3)
	ll.AddAtHead(2)
	ll.AddAtHead(1)

	linkedlist.ReverseListNaive(ll.GetHead())

}
