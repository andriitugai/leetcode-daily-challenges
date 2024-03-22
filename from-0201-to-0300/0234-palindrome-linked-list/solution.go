/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    vals := []int{}
    p := head
    for p != nil {
        vals = append(vals, p.Val)
        p = p.Next
    }
    left, right := 0, len(vals) - 1
    for left < right {
        if vals[left] != vals[right] {
            return false
        }
        left += 1
        right -= 1
    }
    return true
}