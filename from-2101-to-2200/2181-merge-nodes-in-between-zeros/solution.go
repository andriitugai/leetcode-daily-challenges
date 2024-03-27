/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeNodes(head *ListNode) *ListNode {
    curSum := 0
    p, pMerged := head, head
    for true {
        curSum += p.Val
        p = p.Next
        if p == nil {
            break
        }
        if p.Val == 0 {
            pMerged.Val = curSum
            if p.Next == nil {
                pMerged.Next = nil
                break
            }
            curSum = 0
            pMerged.Next = p
            pMerged = pMerged.Next
        }
    }
    return head
}