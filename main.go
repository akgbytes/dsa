package main

import (
	"fmt"

	"github.com/akgbytes/dsa/linkedlist"
)

func main() {

	ll := linkedlist.LinkedList{}

	ll.AddAtHead(1)
	ll.AddAtHead(2)
	ll.AddAtHead(3)
	ll.AddAtHead(3)
	ll.AddAtHead(2)
	ll.AddAtHead(1)

	ll.Print()

	fmt.Println(linkedlist.IsPalindrome(ll.GetHead()))

}
