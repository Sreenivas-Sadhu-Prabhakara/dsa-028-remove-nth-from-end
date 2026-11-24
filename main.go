package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// RemoveNthFromEnd deletes the n-th node counting from the end (n=1 is the
// last node) in one pass, using two pointers kept n apart. If n exceeds the
// list length, the list is returned unchanged.
func RemoveNthFromEnd(head *Node, n int) *Node {
	dummy := &Node{Next: head}
	fast := dummy
	for i := 0; i < n; i++ {
		if fast.Next == nil {
			return head // n longer than the list: nothing to remove
		}
		fast = fast.Next
	}
	slow := dummy
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	fmt.Println(toSlice(RemoveNthFromEnd(fromSlice([]int{1, 2, 3, 4, 5}), 2)))
}
