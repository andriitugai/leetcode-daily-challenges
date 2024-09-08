/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func splitListToParts(head *ListNode, k int) []*ListNode {
    // Get the length of the list
    p := head
    length := 0
    for p != nil {
        p = p.Next
        length += 1
    }
    result := make([]*ListNode, 0)
    for length > 0 || k > 0 {
        if length == 0 {
            for i:=0; i < k; i++ {
                result = append(result, nil)
            }
            break
        }
        chunk := length / k
        if length % k != 0 {
            chunk += 1
        }
        p = head
        for i := 1; i < chunk; i++ {
            p = p.Next
        }
        result = append(result, head)
        head = p.Next
        p.Next = nil

        length -= chunk
        k -= 1
    }
    return result   
}