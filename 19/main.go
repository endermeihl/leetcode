package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := head
	peg := head
	len := 0
	for head != nil {
		head = head.Next
		len++
	}
	if len == 1 && n == 1 {
		return nil
	}
	if len == n {
		return result.Next
	}
	for i := 1; i < len-n; i++ {
		peg = peg.Next
	}
	if peg.Next != nil {
		peg.Next = peg.Next.Next
	} else {
		peg.Next = nil
	}

	return result

}

func removeNthFromEnd_p(head *ListNode, n int) *ListNode {
	result := &ListNode{0, head}
	peg := result
	len := 0
	for head != nil {
		head = head.Next
		len++
	}
	for i := 1; i < len-n; i++ {
		peg = peg.Next
	}
	peg.Next = peg.Next.Next

	return result.Next
}

func removeNthFromEnd_double(head *ListNode, n int) *ListNode {
	result := &ListNode{0, head}
	first := result
	secode := result
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		secode = secode.Next
	}
	secode.Next = secode.Next.Next
	return result.Next
}
