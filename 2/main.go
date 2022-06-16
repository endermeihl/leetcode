package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = new(ListNode)
	result := head
	add := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + add
		mod := sum % 10
		add = sum / 10
		var node = ListNode{Val: mod}
		head.Next = &node
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + add
		mod := sum % 10
		add = sum / 10
		var node = ListNode{Val: mod}
		head.Next = &node
		head = head.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + add
		mod := sum % 10
		add = sum / 10
		var node = ListNode{Val: mod}
		head.Next = &node
		head = head.Next
		l2 = l2.Next
	}
	if add != 0 {
		var node = ListNode{Val: add}
		head.Next = &node
	}

	return result.Next
}

func main() {

}
