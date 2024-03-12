/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeZeroSumSublists(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    prefixSum := 0
    sumMap := make(map[int]*ListNode)
    sumMap[0] = dummy
    for node := dummy; node != nil; node = node.Next {
        prefixSum += node.Val
        sumMap[prefixSum] = node
    }
    prefixSum = 0
    for node := dummy; node != nil; node = node.Next {
        prefixSum += node.Val
        node.Next = sumMap[prefixSum].Next
    }
    return dummy.Next
}