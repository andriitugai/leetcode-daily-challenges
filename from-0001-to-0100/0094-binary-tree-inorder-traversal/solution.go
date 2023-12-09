/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type stackTreeNode []*TreeNode

 func (s *stackTreeNode) push(node *TreeNode) {
	 *s = append(*s, node)
 }
 
 func(s *stackTreeNode) pop() *TreeNode {
	 t := (*s)[len(*s) - 1]
	 *s = (*s)[:len(*s) - 1]
	 return t
 }
 
 func (s stackTreeNode) isEmpty() bool {
	 return len(s) == 0
 }
 
 func inorderTraversal(root *TreeNode) []int {
	 if root == nil {
		 return nil
	 }
	 result := []int{}
	 curr := root
	 
	 s := stackTreeNode{}
 
	 for curr != nil || !s.isEmpty() {
		 for curr != nil {
			 s.push(curr)
			 curr = curr.Left
		 }
		 curr = s.pop()
		 result = append(result, curr.Val)
		 curr = curr.Right
	 }
	 return result
 }