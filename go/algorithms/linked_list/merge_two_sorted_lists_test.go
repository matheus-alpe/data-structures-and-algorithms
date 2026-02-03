package linkedlist

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "Both lists are empty",
			l1:   nil,
			l2:   nil,
			want: nil,
		},
		{
			name: "One list is empty",
			l1:   &ListNode{Val: 1, Next: nil},
			l2:   nil,
			want: &ListNode{Val: 1, Next: nil},
		},
		{
			name: "Both lists have elements",
			l1:   &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
			l2:   &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4, Next: nil},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoSortedLists(tt.l1, tt.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}

}
