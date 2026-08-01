/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    // need to store index and visited node
	hashMap := make(map[*ListNode]bool)
	// i := 0
	// hashMap
	if head == nil{
		return false
	}
	for head.Next != nil{
		hashMap[head] = true
		// i++
		head = head.Next
		if _, exist := hashMap[head]; exist {
			return true
		}
	} 
	return false
}
