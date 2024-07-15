/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func modifiedList(nums []int, head *ListNode) *ListNode {
    exclude := make(map[int]bool)
    for _, num := range nums {
        exclude[num] = true
    }
    dummy := &ListNode{-1, head}
    curr := dummy
    for head != nil {
        if !exclude[head.Val] {
            curr.Next = head
            curr = curr.Next
        }
        head = head.Next
    }
    curr.Next = nil
    return dummy.Next
}