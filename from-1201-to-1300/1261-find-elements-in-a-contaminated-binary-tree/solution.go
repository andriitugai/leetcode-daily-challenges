/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type FindElements struct {
    vals map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    vals := make(map[int]bool)
    var dfs func(node *TreeNode, nodeVal int) 
    dfs = func(node *TreeNode, nodeVal int){
        if node == nil {
            return
        }
        vals[nodeVal] = true
        dfs(node.Left, nodeVal * 2 + 1)
        dfs(node.Right, nodeVal * 2 + 2)
        return
    }
    dfs(root, 0)
    return FindElements{vals: vals}
}


func (this *FindElements) Find(target int) bool {
    _, ok := this.vals[target]
    return ok
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */