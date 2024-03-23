/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    idx := 0
    p := head
    nodes := make(map[int]*ListNode)
    for p != nil {
        nodes[idx] = p
        p = p.Next
        idx += 1
    }

    left, right := 1, idx - 1
    p = head
    for left < right {
        p.Next = nodes[right]
        p = p.Next
        p.Next = nodes[left]
        p = p.Next
        right -= 1
        left += 1
    }

    if left == right {
        p.Next = nodes[left]
        p = p.Next
    }
    p.Next = nil
}