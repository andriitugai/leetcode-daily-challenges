# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumOfLeftLeaves(self, root: Optional[TreeNode]) -> int:
        def dfs(node, is_left: bool)-> int:
            if is_left and node.left is None and node.right is None:
                return node.val
            
            ans = 0
            if node.left:
                ans += dfs(node.left, True)
            if node.right:
                ans += dfs(node.right, False)
                
            return ans
        
        return dfs(root, False)