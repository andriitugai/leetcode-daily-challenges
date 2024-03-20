/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0, Next: list1}
    p1 := dummy
    c1 := 0

    for c1 != a {
        p1 = p1.Next
        c1 += 1
    }
    p2 := p1
    for c1 != b + 1 {
        p1 = p1.Next
        c1 += 1
    }
    p1 = p1.Next
    p2.Next = list2
    for p2.Next != nil {
        p2 = p2.Next
    }
    p2.Next = p1

    return dummy.Next
}