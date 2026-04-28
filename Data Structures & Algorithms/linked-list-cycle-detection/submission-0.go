/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	nodeSet := map[*ListNode]bool{}
	curr := head
	for curr != nil {
		if nodeSet[curr.Next] { return true }
		nodeSet[curr] = true
		curr = curr.Next
	}
	return false
}
