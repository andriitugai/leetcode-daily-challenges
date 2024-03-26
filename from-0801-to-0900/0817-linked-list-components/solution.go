/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func numComponents(head *ListNode, nums []int) int {
    inNums := make(map[int]bool)
    for _, num := range nums {
        inNums[num] = true
    }

    nGroups := 0
    inChain := false
    p := head
    for p != nil {
        if inNums[p.Val] {
            if !inChain {
                inChain = true
                nGroups += 1
            }
        } else {
            if inChain {
                inChain = false
            }
        }
        p = p.Next
    }
    return nGroups
}