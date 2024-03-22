/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func getDirections(root *TreeNode, startValue int, destValue int) string {
    
    pathToSrc := []string{}
    pathToDst := []string{}
    dfs(root, startValue, &pathToSrc)
    dfs(root, destValue, &pathToDst)

    reverse(pathToSrc)
    reverse(pathToDst)

    i := 0
    for i < len(pathToSrc) && i < len(pathToDst) {
        if pathToSrc[i] == pathToDst[i] {
            i += 1
        } else {
            break
        }
    }
    pathToSrc = pathToSrc[i:]
    pathToDst = pathToDst[i:]
    
    for i := 0; i < len(pathToSrc); i ++ {
        pathToSrc[i] = "U"
    }
    
    return strings.Join(pathToSrc, "") + strings.Join(pathToDst, "")
}

func dfs(node *TreeNode, target int, path *[]string) bool {
    if node == nil {
        return false
    }
    if node.Val == target {
        return true
    }
    if node.Left != nil {
        if dfs(node.Left, target, path) {
            *path = append(*path, "L")
            return true
        }
    }
    if node.Right != nil {
        if dfs(node.Right, target, path) {
            *path = append(*path, "R")
            return true
        }
    }
    return false
}

func reverse(lst []string){
    left, right := 0, len(lst) - 1
    for left < right {
        lst[left], lst[right] = lst[right], lst[left]
        left += 1
        right -= 1
    }
}