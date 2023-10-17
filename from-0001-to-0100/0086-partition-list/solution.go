/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }

    headLeft, headRight := &ListNode{}, &ListNode{}
    left, right := headLeft, headRight
    for head != nil {
        if head.Val < x {
            left.Next = head
            left = left.Next
        } else {
            right.Next = head
            right = right.Next
        }
        head = head.Next
    }
    left.Next, right.Next = nil, nil
    if headRight.Next != nil {
        left.Next = headRight.Next
    }
    return headLeft.Next   
}