package problem160

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1, l2 := headA, headB
	for {
		if l1 == l2 {
			return l1
		}
		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			return nil
		}
		if l1 == nil {
			l1 = headB
		}
		if l2 == nil {
			l2 = headA
		}
	}
}
