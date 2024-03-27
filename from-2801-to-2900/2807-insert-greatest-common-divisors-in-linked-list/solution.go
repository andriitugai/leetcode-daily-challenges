/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    p := head
    var newVal int
    var newNode *ListNode
    for p.Next != nil {
        newVal = gcd(p.Val, p.Next.Val)
        newNode = &ListNode{Val: newVal, Next: p.Next}
        p.Next = newNode
        p = p.Next.Next
    }
    return head
}

func gcd(a, b int) int {
    for a != b {
        if a > b {
            a -= b
        } else {
            b -= a
        }
    }
    return a
}