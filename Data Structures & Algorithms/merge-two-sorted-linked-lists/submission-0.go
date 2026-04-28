/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// base case
	if list1 == nil { return list2 }
	if list2 == nil { return list1 }

	// get the first item for the list
	var list *ListNode
	if list1.Val < list2.Val {
		list = list1
		list1 = list1.Next
	} else {
		list = list2
		list2 = list2.Next
	}
	result := list

	// iterate over all until one list  got to nil
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list.Next = list1
			list, list1 = list.Next, list1.Next
		} else {
			list.Next = list2
			list, list2 = list.Next, list2.Next
		}
	}

	// get the remaining, we will get into one of these
	if list1 != nil { list.Next = list1 }
	if list2 != nil { list.Next = list2 }
    
	return result
}
