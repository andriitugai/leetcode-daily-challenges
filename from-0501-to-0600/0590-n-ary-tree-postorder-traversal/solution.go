/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func postorder(root *Node) []int {
    result := []int{}
    var dfs func(node *Node)
    dfs = func(node *Node){
        if node == nil {
            return
        }
        for _, c := range node.Children {
            dfs(c)
        }
        result = append(result, node.Val)
    }
    dfs(root)
    return result
}