package main

import "fmt"

type Nodes struct {
	Val  int
	Next *Nodes
}

func (head *Nodes) Stagger() {
	current := head
	for head != nil {
		if head.Next == nil {
			fmt.Println(head)
			return
		}
		if head.Next.Next == nil {
			break
		}
		step1 := head.Next
		step2 := head.Next.Next
		step3 := head.Next.Next.Next
		step1.Next = step3
		step2.Next = step1
		head.Next = step2
		head = step1
		continue
	}
	fmt.Println(current)
}

//Reverse a Linkedlist
type SinglyLinkedListNode struct {
	Val  int
	next *SinglyLinkedListNode
}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	var previous *SinglyLinkedListNode
	for llist != nil {
		next := llist.next
		llist.next = previous
		previous = llist
		llist = next
	}
	return previous

}
