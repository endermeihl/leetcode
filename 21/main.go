package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head = new(ListNode)
	result := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			var node = ListNode{Val: list2.Val}
			head.Next = &node
			list2 = list2.Next
		} else {
			Node := ListNode{list1.Val, nil}
			head.Next = &Node
			list1 = list1.Next
		}
		head = head.Next
	}
	for list1 != nil {
		Node := ListNode{list1.Val, nil}
		head.Next = &Node
		list1 = list1.Next
		head = head.Next
	}

	for list2 != nil {
		Node := ListNode{list2.Val, nil}
		head.Next = &Node
		list2 = list2.Next
		head = head.Next
	}

	return result.Next
}
