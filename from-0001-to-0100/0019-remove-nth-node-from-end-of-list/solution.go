/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cnt := 1
    var curr, prev *ListNode = head, head
    curr = head
    for curr.Next != nil {
        cnt += 1
        curr = curr.Next
        if cnt > n + 1{
            prev = prev.Next
        }
    }
    if n == cnt {
        return head.Next
    }
    if prev.Next != nil {
        prev.Next = prev.Next.Next
    }
    return head
}