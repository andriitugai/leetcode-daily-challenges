/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func recoverFromPreorder(traversal string) *TreeNode {
    idx, n := 0, len(traversal)
    stack := []*TreeNode{}
    num := ""

    for idx < n {
        lvl := 0
        for idx < n && traversal[idx] == '-' {
            idx += 1
            lvl += 1
        }
        for len(stack) > lvl {
            stack = stack[:len(stack) - 1]
        }
        for idx < n && traversal[idx] != '-' {
            num += string(traversal[idx])
            idx += 1
        }
        val, _ := strconv.Atoi(num)
        node := &TreeNode{Val: val}

        if len(stack) > 0 && stack[len(stack) - 1].Left == nil {
            stack[len(stack) - 1].Left = node
        } else if len(stack) > 0 {
            stack[len(stack) - 1].Right = node
        }

        stack = append(stack, node)
        num = ""
    }
    return stack[0]
}