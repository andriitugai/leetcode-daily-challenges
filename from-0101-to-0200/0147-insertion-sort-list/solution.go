/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func insertionSortList(head *ListNode) *ListNode {
    dummy := &ListNode{Val: -10000, Next: head}
    var nodeToInsert, pPlaceToInsert, pProgress *ListNode = dummy, dummy, head
    for pProgress != nil && pProgress.Next != nil {
        if pProgress.Val > pProgress.Next.Val {
            nodeToInsert = pProgress.Next
            pPlaceToInsert = dummy
            for pPlaceToInsert.Next.Val < nodeToInsert.Val {
                pPlaceToInsert = pPlaceToInsert.Next
            }
            pProgress.Next = nodeToInsert.Next
            nodeToInsert.Next = pPlaceToInsert.Next
            pPlaceToInsert.Next = nodeToInsert
        } else {
            pProgress = pProgress.Next
        }
    }
    return dummy.Next
}