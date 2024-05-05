/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(node *ListNode) {
    for true {
        node.Val = node.Next.Val
        if node.Next.Next != nil {
            node = node.Next
        } else {
            node.Next = nil
            break
        }
    }
}