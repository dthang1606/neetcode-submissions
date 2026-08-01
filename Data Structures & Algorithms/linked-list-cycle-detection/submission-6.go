/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    // need to store index and visited node 
	// floyd algo with fast/slow pointer
	// fast pointer will eventually meet the slow pointer
	if head == nil {
    	return false
    }
	slow := head
	fast := head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false

}

// func hasCycle(head *ListNode) bool {
//     // need to store index and visited node 
// 	// Time complexity: O(n)
// 	// Space complexity: O(n)
// 	hashMap := make(map[*ListNode]bool)
// 	if head == nil{
// 		return false
// 	}
// 	for head.Next != nil{
// 		hashMap[head] = true
// 		head = head.Next
// 		if _, exist := hashMap[head]; exist {
// 			return true
// 		}
// 	} 
// 	return false
// }
