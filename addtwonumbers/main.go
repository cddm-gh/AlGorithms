package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {
	node := Node{}
	node.Val = val
	if l.len == 0 {
		l.head = &node
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.Next == nil {
			ptr.Next = &node
			l.len++
			return
		}
		ptr = ptr.Next
	}
}

func (l *LinkedList) InserSum(val1, val2 int) int {
	node := Node{}
	fmt.Printf("Adding sum %d value to node\n", val1+val2)
	sum := val1 + val2
	var add int
	var remainder int
	if sum == 10 {
		node.Val = 0
		remainder = 1
	} else if sum > 10 {
		add = sum - 10
		remainder = int(math.Floor(float64(val1) / float64(val2)))
		node.Val = add
	} else {
		node.Val = sum
	}
	if l.len == 0 {
		l.head = &node
		l.len++
		return 0
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.Next == nil {
			ptr.Next = &node
			l.len++
			return 0
		}
		ptr = ptr.Next
	}
	return remainder
}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("The list is empty")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.Next
	}
}

func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.Val == val {
			return i
		}
		ptr = ptr.Next
	}
	return -1
}

func main() {
	l1 := LinkedList{}
	l2 := LinkedList{}

	l1.Insert(2)
	l1.Insert(4)
	l1.Insert(6)
	l2.Insert(3)
	l2.Insert(5)
	l2.Insert(4)

	l1.Print()
	l2.Print()

	result := addTwoNumbers(&l1, &l2)

	fmt.Println("The result list is: ")
	result.Print()
}

func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	ptr := l1.head
	ptr2 := l2.head
	lr := LinkedList{}

	for ptr != nil {
		if ptr2 != nil {
			lr.InserSum(ptr.Val, ptr2.Val)
			ptr = ptr.Next
			ptr2 = ptr2.Next
		}
	}

	return &lr
}
