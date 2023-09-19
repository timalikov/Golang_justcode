package task2

import "testing"

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		l1       *ListNode
		l2       *ListNode
		expected []int
	}{
		{
			l1:       &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			l2:       &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			l1:       &ListNode{1, nil},
			l2:       &ListNode{2, nil},
			expected: []int{1, 2},
		},
		{
			l1:       nil,
			l2:       &ListNode{1, &ListNode{3, nil}},
			expected: []int{1, 3},
		},
		{
			l1:       &ListNode{1, &ListNode{2, nil}},
			l2:       nil,
			expected: []int{1, 2},
		},
		{
			l1:       nil,
			l2:       nil,
			expected: []int{},
		},
	}

	for i, tc := range testCases {
		result := mergeTwoLists(tc.l1, tc.l2)
		for j, val := range tc.expected {
			if result == nil {
				t.Fatalf("Test case %d failed: expected %v but got nil", i+1, tc.expected)
			}
			if result.Val != val {
				t.Fatalf("Test case %d failed: expected %d at position %d but got %d", i+1, val, j, result.Val)
			}
			result = result.Next
		}
		if result != nil {
			t.Fatalf("Test case %d failed: result list has more elements than expected", i+1)
		}
	}
}
