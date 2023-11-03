package task2

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(n+m), where n is the length of the first linked list and m is the length of the second linked list.

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newNode := &ListNode{}
	current := newNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
		list1 = list1.Next
	}

	if list2 != nil {
		current.Next = list2
		list2 = list2.Next
	}

	return newNode.Next

}
