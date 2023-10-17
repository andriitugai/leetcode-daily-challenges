/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseBetween(head *ListNode, left int, right int) *ListNode {
    p_slow, p_fast := head, head
    pos := 1

    for pos < left {
        p_slow = p_slow.Next
        p_fast = p_fast.Next
        pos += 1
    }

    stack := []int{}
    for pos <= right {
        stack = append(stack, p_fast.Val)
        p_fast = p_fast.Next
        pos += 1
    }

    for i := len(stack)-1; i >= 0; i-- {
        p_slow.Val = stack[i]
        p_slow = p_slow.Next
    }
    return head
}