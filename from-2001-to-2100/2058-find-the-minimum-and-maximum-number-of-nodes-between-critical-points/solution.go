/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func nodesBetweenCriticalPoints(head *ListNode) []int {
    cpt := []int{}
    prev := head.Val
    p := head.Next
    idx := 2

    for p.Next != nil {
        if prev < p.Val && p.Next.Val < p.Val || prev > p.Val && p.Next.Val > p.Val {
            cpt = append(cpt, idx)
        }
        prev = p.Val
        p = p.Next
        idx += 1
    }
    
    if len(cpt) < 2 {
        return []int{-1, -1}
    }
    minDist := 1000000
    var dist int
    for i := 1; i < len(cpt); i ++ {
        dist = cpt[i] - cpt[i - 1]
        if dist < minDist {
            minDist = dist
        }
    }
    return []int{minDist, cpt[len(cpt) - 1] - cpt[0]}
}