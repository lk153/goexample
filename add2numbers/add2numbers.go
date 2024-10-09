package add2numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	return addNumbers(l1, l2, result)
}

func addNumbers(l1 *ListNode, l2 *ListNode, result *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return result
	}

	if l1 != nil && l1.Val != 0 {
		result.Val += l1.Val
	}

	if l2 != nil && l2.Val != 0 {
		result.Val += l2.Val
	}

	if result.Val > 9 {
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}

		if l2.Next == nil {
			l2.Next = &ListNode{}
		}

		l1.Next.Val += result.Val / 10
		result.Val = result.Val % 10
		result.Next = &ListNode{}
		return addNumbers(l1.Next, l2.Next, result.Next)
	}

	switch {
	case l1 == nil && l2.Next == nil:
		return result
	case l2 == nil && l1.Next == nil:
		return result
	case l1 != nil && l1.Next == nil && l2 != nil && l2.Next == nil:
		return result
	}

	result.Next = &ListNode{}
	switch {
	case l1 == nil:
		return addNumbers(nil, l2.Next, result.Next)
	case l2 == nil:
		return addNumbers(l1.Next, nil, result.Next)
	default:
		return addNumbers(l1.Next, l2.Next, result.Next)
	}
}

// func main() {
// 	type testcase struct {
// 		input struct {
// 			l1 *ListNode
// 			l2 *ListNode
// 		}
// 		ouput struct {
// 			result *ListNode
// 		}
// 	}
// 	testcases := map[string]testcase{
// 		"case-1": {
// 			input: struct {
// 				l1 *ListNode
// 				l2 *ListNode
// 			}{
// 				l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}},
// 				l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}},
// 			},
// 			ouput: struct {
// 				result *ListNode
// 			}{
// 				result: &ListNode{},
// 			},
// 		},
// 	}

// 	for _, tc := range testcases {
// 		result := addTwoNumbers(tc.input.l1, tc.input.l2)
// 		fmt.Printf("%#v", result)
// 	}
// }
