/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type nodeInfo struct {
    idx int
    val int
}

func nextLargerNodes(head *ListNode) []int {
    result := []int{}
    p := head
    stack := []*nodeInfo{}
    cnt := 0

    for p != nil {
        result = append(result, 0)
        for len(stack) > 0  && p.Val > stack[len(stack) - 1].val {
            currIdx, _ := stack[len(stack) - 1].idx, stack[len(stack) - 1].val
            stack = stack[:len(stack) - 1]
            result[currIdx] = p.Val
        }
        stack = append(stack, &nodeInfo{idx: cnt, val: p.Val})
        cnt += 1
        p = p.Next
    }
    return result
}