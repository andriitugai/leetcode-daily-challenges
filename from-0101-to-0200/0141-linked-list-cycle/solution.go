/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    fast, slow := head, head
    for fast.Next != nil {
        fast = fast.Next
        if fast == slow {
            return true
        }
        fast = fast.Next
        if fast == nil {
            break
        }
        if fast == slow {
            return true
        }
        slow = slow.Next
    }
    return false
}