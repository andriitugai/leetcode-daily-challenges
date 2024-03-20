/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func smallestFromLeaf(root *TreeNode) string {
    allStrings := []string{}
    var dfs func(node *TreeNode, curString []string)
    dfs = func(node *TreeNode, curString []string){
        curString = append([]string{string('a' + node.Val)}, curString...)
        if node.Left == nil && node.Right == nil {
            allStrings = append(allStrings, strings.Join(curString, ""))
            return
        }
        if node.Left != nil {
            dfs(node.Left, curString)
        }
        if node.Right != nil {
            dfs(node.Right, curString)
        }
        return
    }
    dfs(root, []string{})
    sort.Strings(allStrings)
    return allStrings[0]
}