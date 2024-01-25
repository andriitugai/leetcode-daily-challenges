# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pseudoPalindromicPaths (self, root: Optional[TreeNode]) -> int:
        ans = 0
        def dfs(node, seen):
            nonlocal ans
            if node.val in seen:
                seen.remove(node.val)
            else:
                seen.add(node.val)
                
            if node.left:
                dfs(node.left, seen.copy())
            if node.right:
                dfs(node.right, seen.copy())
                
            if not node.left and not node.right:
                if len(seen) < 2:
                    ans += 1
                    
        dfs(root, set())            
        return ans
    