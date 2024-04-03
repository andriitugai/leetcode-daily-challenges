/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func doubleIt(head *ListNode) *ListNode {
    tail := reverseLinkedList(head)
    var dbl, tr int = 0, 0
    var prev *ListNode
    p := tail
    for p != nil {
        dbl = p.Val * 2 + tr
        p.Val = dbl % 10
        tr = dbl / 10
        prev = p
        p = p.Next
    }
    if tr > 0 {
        prev.Next = &ListNode{Val: tr}
        prev = prev.Next
    }
    prev.Next = nil
    return reverseLinkedList(tail)
}

func reverseLinkedList(head *ListNode) *ListNode {
    var prev, next *ListNode
    curr := head
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}