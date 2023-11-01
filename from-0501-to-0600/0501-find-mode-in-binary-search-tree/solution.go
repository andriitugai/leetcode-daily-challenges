/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func dfs(node *TreeNode, cnt *map[int]int) {
    if node == nil {
            return
        }
        (*cnt)[node.Val] += 1
        dfs(node.Left, cnt)
        dfs(node.Right, cnt)
}

func findMode(root *TreeNode) []int {
    cnt := make(map[int]int)
    
    dfs(root, &cnt)
    mode := -1000001
    for _, val := range cnt {
        if val > mode {
            mode = val
        }
    }
    result := []int{}
    for key, val := range cnt {
        if val == mode {
            result = append(result, key)
        }
    }
    return result
}