/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
        if fast.Next == nil {
            break
        }
        fast = fast.Next
    }
    return slow
}