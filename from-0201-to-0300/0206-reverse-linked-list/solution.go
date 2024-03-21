/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := head
    var neck *ListNode
    for head.Next != nil {
        neck = head.Next
        head.Next = neck.Next
        neck.Next = newHead
        newHead = neck
    }
    return newHead
}