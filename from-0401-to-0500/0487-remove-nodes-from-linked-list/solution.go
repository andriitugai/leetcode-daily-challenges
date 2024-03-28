/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNodes(head *ListNode) *ListNode {
    var rmNodes func(node *ListNode) *ListNode
    rmNodes = func(node *ListNode) *ListNode {
        if node == nil {
            return nil
        }
        node.Next = rmNodes(node.Next)
        if node.Next != nil && node.Val < node.Next.Val {
            return node.Next
        }
        return node
    }
    return rmNodes(head)
}